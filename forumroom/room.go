package forumroom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type (
	Room struct {
		sync.RWMutex
		clients   map[*Client]bool
		broadcast chan []byte
		join      chan *Client
		leave     chan *Client
	}
	Client struct {
		room *Room
		conn *websocket.Conn
		send chan []byte
	}
	MESSAGE map[string]interface{}
)

var Mid uint64

//New forum Room
func NewRoom() *Room {
	return &Room{
		broadcast: make(chan []byte),
		join:      make(chan *Client),
		leave:     make(chan *Client),
		clients:   make(map[*Client]bool),
	}
}

func (r *Room) AddClient(c *Client) {
	r.Lock()
	defer r.Unlock()
	r.clients[c] = true
}

func (r *Room) RemoveClient(c *Client) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.clients[c]; ok {
		delete(r.clients, c)
		close(c.send)
	}
	return nil
}

func (r *Room) TotalNums() int {
	r.RLock()
	defer r.RUnlock()
	return len(r.clients)
}

func (r *Room) BroadcastMsg(message []byte) {
	r.Lock()
	defer r.Unlock()
	var resMsg MESSAGE
	json.Unmarshal(message, &resMsg)

	msg, _ := json.Marshal(MESSAGE{
		"mid":     atomic.AddUint64(&Mid, 1),
		"time":    time.Now().UnixNano() / 1000000,
		"content": resMsg,
	})

	for client := range r.clients {
		select {
		case client.send <- msg:
		default:
			close(client.send)
			r.RemoveClient(client)
		}
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.AddClient(client)
		case client := <-r.leave:
			r.RemoveClient(client)
		case message := <-r.broadcast:
			r.BroadcastMsg(message)
		}
	}
}

func (c *Client) ReceiveLoop() {
	defer func() {
		c.room.leave <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Fprintf(gin.DefaultWriter, "c.conn.ReadMessage error:%v", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, []byte{'\n'}, []byte{' '}, -1))
		fmt.Fprintf(gin.DefaultWriter, "Rec from %s msg:%v\n", c.conn.RemoteAddr().String(), string(msg))
		c.room.broadcast <- msg
	}

}

func (c *Client) SendLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			fmt.Fprintf(gin.DefaultWriter, "Send to %s msg:%v\n", c.conn.RemoteAddr().String(), string(msg))
			if err = w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

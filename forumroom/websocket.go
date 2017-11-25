/*
|--------------------------------------------------------------------------
| webSocket
|--------------------------------------------------------------------------
|
| webSocket
|
*/
package forumroom

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/qclaogui/goforum/model"
)

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"forumToken"},
}

func ServerWS(room *Room, w http.ResponseWriter, r *http.Request) {
	wsConn, err := wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(gin.DefaultWriter, "Failed to set webSocket upgrade:%v", err)
	}
	//	验证token
	_, err = CheckWebSocketToken(r)
	if err != nil {
		fmt.Fprintf(gin.DefaultWriter, "WebSocket Token Error:%v", err.Error())
		msg := MESSAGE{
			"error_code": 40000,
			"message":    "Invalid Token",
		}
		wsConn.WriteJSON(msg)
	}

	client := &Client{room: room, conn: wsConn, send: make(chan []byte, 256)}
	client.room.join <- client

	go client.SendLoop()
	go client.ReceiveLoop()

	fmt.Fprintf(gin.DefaultWriter, "room clients num [%d] user\n", len(client.room.clients)+1)
}

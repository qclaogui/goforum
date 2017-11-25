package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/forumroom"
)

func WebSocketGroup(r *gin.Engine) {

	room := forumroom.NewRoom()
	go room.Run()

	r.GET("ws", func(c *gin.Context) {
		forumroom.ServerWS(room, c.Writer, c.Request)
	})
}

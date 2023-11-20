package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

func Connect(c *gin.Context) {
	log.Println("connected ")
	wsConn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ErrorResponse(c.Writer, "upgrade error")
		return
	}

	wsConn.WriteMessage(websocket.TextMessage, []byte("connected"))
}

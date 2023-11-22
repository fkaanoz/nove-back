package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Connect(c *gin.Context) {
	wsConn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ErrorResponse(c.Writer, "upgrade error")
		return
	}

	wsConn.WriteMessage(websocket.TextMessage, []byte("connected"))
}

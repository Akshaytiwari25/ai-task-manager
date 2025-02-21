package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // For development; restrict in production.
	},
}

// HandleWebSocket upgrades the HTTP connection to a WebSocket and echoes messages.
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected via WebSocket")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Println("Received message:", string(msg))

		if err := conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
	fmt.Println("Client disconnected")
}

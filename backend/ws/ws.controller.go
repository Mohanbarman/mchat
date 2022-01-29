package ws

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func WsController(ctx *gin.Context) {
	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed to upgrade connection to websocket",
		})
		return
	}

	defer connection.Close()

	for {
		mt, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", message)
		err = connection.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}

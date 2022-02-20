package ws

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/ws/connection"
	"mchat.com/api/modules/ws/events"
)

type WsController struct {
	DB       *gorm.DB
	Upgrader websocket.Upgrader
	Config   *config.Config
}

func (ctrl *WsController) CreateConnection(manager *connection.ConnStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, err := ctrl.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if err != nil {
			ctx.JSON(400, gin.H{
				"success": false,
				"code":    400,
				"message": "Failed to upgrade connection to websocket",
			})
			return
		}

		defer c.Close()

		context := connection.Context{
			Connection: c,
			DB:         ctrl.DB,
			Config:     ctrl.Config,
		}

		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				break
			}

			wsMessage := events.Event{}
			e := json.Unmarshal(message, &wsMessage)

			if e == nil {
				events.HandleEvent(&wsMessage, &context, manager)
			}
		}
	}
}

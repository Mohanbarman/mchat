package connection

import (
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/models"
)

type Context struct {
	Connection *websocket.Conn
	User       *models.UserModel
	DB         *gorm.DB
	Config     *config.Config
}

func (c *Context) Send(event string, payload map[string]interface{}) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"payload": payload,
	})
}

func (c *Context) SendMessage(event string, payload string) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"payload": payload,
	})
}

func (c *Context) SendErr(event string, code string) {
	c.SendJSON(map[string]interface{}{
		"event": event,
		"code":  code,
	})
}

func (c *Context) SendErrP(event string, code string, payload map[string]interface{}) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"code":    code,
		"payload": payload,
	})
}

func (c *Context) SendJSON(payload map[string]interface{}) {
	c.Connection.WriteJSON(payload)
}

func (c *Context) Close() {
	c.Connection.Close()
}

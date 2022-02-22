package lib

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/models"
)

// responsible for managing the context of websocket
type WsContext struct {
	Connection *websocket.Conn
	User       *models.UserModel
	DB         *gorm.DB
	Config     *config.Config
}

func (c *WsContext) Send(event string, payload interface{}) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"payload": payload,
	})
}

func (c *WsContext) SendMessage(event string, payload string) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"payload": payload,
	})
}

func (c *WsContext) SendErr(event string, code string) {
	c.SendJSON(map[string]interface{}{
		"event": event,
		"code":  code,
	})
}

func (c *WsContext) SendErrP(event string, code string, payload map[string]interface{}) {
	c.SendJSON(map[string]interface{}{
		"event":   event,
		"code":    code,
		"payload": payload,
	})
}

func (c *WsContext) SendJSON(payload map[string]interface{}) {
	c.Connection.WriteJSON(payload)
}

func (c *WsContext) Close() {
	c.Connection.Close()
}

// responsible for storing connected websocket connections
type WsStore struct {
	connections map[uint]*websocket.Conn
}

func (c *WsStore) Set(key uint, con *websocket.Conn) {
	c.connections[key] = con
}

func (c *WsStore) Get(key uint) (*WsContext, error) {
	if con, ok := c.connections[key]; ok {
		ctx := &WsContext{
			Connection: con,
		}
		return ctx, nil
	}
	return nil, errors.New("key not found")
}

func (c *WsStore) Remove(key uint) {
	delete(c.connections, key)
}

func NewWsStore() *WsStore {
	return &WsStore{
		connections: make(map[uint]*websocket.Conn),
	}
}

// helpers
func ValidatePayload(src interface{}, dst interface{}) (err error) {
	mapstructure.Decode(src, dst)
	validate := validator.New()
	err = validate.Struct(dst)
	return
}

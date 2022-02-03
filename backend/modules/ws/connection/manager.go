package connection

import (
	"errors"

	"github.com/gorilla/websocket"
)

type ConnManager struct {
	connections map[uint]*websocket.Conn
}

func (c *ConnManager) Set(key uint, con *websocket.Conn) {
	c.connections[key] = con
}

func (c *ConnManager) Get(key uint) (*Context, error) {
	if con, ok := c.connections[key]; ok {
		ctx := &Context{
			Connection: con,
		}
		return ctx, nil
	}
	return nil, errors.New("key not found")
}

func (c *ConnManager) Remove(key uint) {
	delete(c.connections, key)
}

func NewManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint]*websocket.Conn),
	}
}

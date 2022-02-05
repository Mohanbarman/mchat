package connection

import (
	"errors"

	"github.com/gorilla/websocket"
)

type ConnStore struct {
	connections map[uint]*websocket.Conn
}

func (c *ConnStore) Set(key uint, con *websocket.Conn) {
	c.connections[key] = con
}

func (c *ConnStore) Get(key uint) (*Context, error) {
	if con, ok := c.connections[key]; ok {
		ctx := &Context{
			Connection: con,
		}
		return ctx, nil
	}
	return nil, errors.New("key not found")
}

func (c *ConnStore) Remove(key uint) {
	delete(c.connections, key)
}

func NewManager() *ConnStore {
	return &ConnStore{
		connections: make(map[uint]*websocket.Conn),
	}
}

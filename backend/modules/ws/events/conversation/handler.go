package conversation

import (
	"mchat.com/api/modules/ws/connection"
)

func HandleEvent(method string, payload interface{}, ctx *connection.Context, store *connection.ConnStore) {
	controller := Controller{
		Store: store,
	}

	if ctx.User == nil {
		ctx.SendMessage("auth/error", "UNAUTHORIZED")
	}

	switch method {
	case "send":
		controller.Send(payload, ctx)
	case "read":
		controller.Read(payload, ctx)
	}
}

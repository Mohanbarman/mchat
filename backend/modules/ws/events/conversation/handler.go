package conversation

import (
	"mchat.com/api/modules/ws/connection"
)

func HandleEvent(method string, payload interface{}, context *connection.Context, manager *connection.ConnStore) {
	controller := Controller{
		Manager: manager,
	}

	switch method {
	case "send":
		controller.Send(payload, context)
	case "read":
		controller.Read(payload, context)
	}
}

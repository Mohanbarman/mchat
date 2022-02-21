package events

import (
	"strings"

	"github.com/gin-gonic/gin"
	"mchat.com/api/lib"
	"mchat.com/api/ws/events/auth"
	"mchat.com/api/ws/events/conversation"
)

type Event struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func HandleEvent(e *Event, ctx *lib.WsContext, manager *lib.WsStore) {
	action := strings.Split(e.Action, "/")

	if len(action) < 2 {
		ctx.SendJSON(gin.H{
			"event": "error",
			"code":  "BAD_REQUEST",
		})
		return
	}

	group := action[0]
	method := action[1]
	payload := e.Payload

	switch group {
	case "conversation":
		conversation.HandleEvent(method, payload, ctx, manager)
	case "auth":
		auth.HandleEvent(method, payload, ctx, manager)
	}
}

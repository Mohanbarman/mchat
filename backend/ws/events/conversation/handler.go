package conversation

import "mchat.com/api/lib"

func HandleEvent(method string, payload map[string]interface{}, ctx *lib.WsContext, store *lib.WsStore) {
	controller := Controller{
		Store: store,
	}

	if ctx.User == nil {
		ctx.SendMessage("auth/error", "UNAUTHORIZED")
		return
	}

	switch method {
	case "send":
		controller.Send(payload, ctx)
	case "read":
		controller.Read(payload, ctx)
	case "typing":
		controller.Typing(payload, ctx)
	}
}

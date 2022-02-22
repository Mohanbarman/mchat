package auth

import (
	"mchat.com/api/lib"
)

func HandleEvent(method string, payload map[string]interface{}, ctx *lib.WsContext, manager *lib.WsStore) {
	controller := Controller{
		Store: manager,
		Jwt: &lib.Jwt{
			Config: &ctx.Config.Jwt,
		},
	}

	switch method {
	case "login":
		controller.Login(payload, ctx)
	}
}

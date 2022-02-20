package auth

import (
	"mchat.com/api/lib/jwt"
	"mchat.com/api/modules/ws/connection"
)

func HandleEvent(method string, payload interface{}, ctx *connection.Context, manager *connection.ConnStore) {
	controller := Controller{
		Store: manager,
		Jwt: &jwt.JwtService{
			Config: &ctx.Config.Jwt,
		},
	}

	switch method {
	case "login":
		controller.Login(payload, ctx)
	}
}

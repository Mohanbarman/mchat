package auth

import (
	"github.com/mitchellh/mapstructure"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/models"
	"mchat.com/api/modules/ws/connection"
)

type Controller struct {
	Store *connection.ConnStore
	Jwt   *jwt.JwtService
}

func (c *Controller) Login(payload interface{}, ctx *connection.Context) {
	dto := LoginDTO{}
	mapstructure.Decode(payload, &dto)

	token := dto.Token

	sub, err := c.Jwt.ParseToken(token, jwt.AccessToken)

	if err != nil {
		ctx.SendErr("auth/errors", "INVALID_TOKEN")
		return
	}

	user := &models.UserModel{}
	if r := ctx.DB.First(user, &models.UserModel{UUID: sub}); r.RowsAffected < 1 {
		ctx.SendErr("auth/errors", "INVALID_TOKEN")
		return
	}

	c.Store.Set(user.ID, ctx.Connection)
	ctx.User = user
	ctx.SendMessage("auth/success", "success")
}

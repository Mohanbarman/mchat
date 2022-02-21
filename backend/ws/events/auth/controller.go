package auth

import (
	"github.com/mitchellh/mapstructure"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Controller struct {
	Store *lib.WsStore
	Jwt   *lib.Jwt
}

func (c *Controller) Login(payload interface{}, ctx *lib.WsContext) {
	dto := LoginDTO{}
	mapstructure.Decode(payload, &dto)

	token := dto.Token

	sub, err := c.Jwt.ParseToken(token, lib.AccessToken)

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

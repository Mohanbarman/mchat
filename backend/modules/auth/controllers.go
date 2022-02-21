package auth

import (
	"github.com/gin-gonic/gin"
	"mchat.com/api/lib"
	"mchat.com/api/models"
	"mchat.com/api/validation"
)

type Controller struct {
	Service *Service
}

func (ctrl *Controller) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &LoginDto{}
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}

		data, err := ctrl.Service.Login(dto)

		if err != nil {
			HttpErrors[err.Code].Send(c)
			return
		}

		lib.HttpResponse(200).Data(data).Send(c)
	}
}

func (ctrl *Controller) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.UserModel)
		lib.HttpResponse(200).Data(user.Transform()).Send(c)
	}
}

func (ctrl *Controller) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &RegisterDto{}
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}
		result, err := ctrl.Service.Register(dto)

		if err != nil {
			HttpErrors[err.Code].Send(c)
			return
		}

		lib.HttpResponse(200).Data(result).Send(c)
	}
}

func (ctrl *Controller) SendResetPasswordMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &ResetPasswordDTO{}
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}
		go ctrl.Service.SendResetPasswordMail(dto)
		HttpSuccess[ResetPassEmailSentSucccess].Send(c)
	}
}

func (ctrl *Controller) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &ResetPasswordChangeDTO{}
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}
		_, err := ctrl.Service.ResetPassword(dto)

		if err != nil {
			HttpErrors[err.Code].Send(c)
			return
		}

		HttpSuccess[PassChangedSuccess].Send(c)
	}
}

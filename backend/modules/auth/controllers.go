package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/models"
)

type AuthController struct {
	Config     *config.Config
	DB         *gorm.DB
	Service    *AuthService
	JwtService *jwt.JwtService
	OtpSmtp    *lib.MailClient
	Redis      *lib.RedisClient
}

func (ctrl *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*LoginDto)

		data, err := ctrl.Service.Login(dto, ctrl.JwtService)

		if err != nil {
			HttpErrors[err.Code].Send(c)
		}

		lib.HttpResponse(200).Data(data).Send(c)
	}
}

func (ctrl *AuthController) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.UserModel)
		lib.HttpResponse(200).Data(user.Transform()).Send(c)
	}
}

func (ctrl *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*RegisterDto)
		result, err := ctrl.Service.Register(dto)

		if err != nil {
			HttpErrors[err.Code].Send(c)
		}

		lib.HttpResponse(200).Data(result).Send(c)
	}
}

func (ctrl *AuthController) SendResetPasswordMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*ResetPasswordDTO)
		go ctrl.Service.SendResetPasswordMail(dto, ctrl.OtpSmtp, ctrl.Redis)
		HttpSuccess[ResetPassEmailSentSucccess].Send(c)
	}
}

func (ctrl *AuthController) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*ResetPasswordChangeDTO)
		_, err := ctrl.Service.ResetPassword(dto, ctrl.Redis)

		if err != nil {
			HttpErrors[err.Code].Send(c)
			return
		}

		HttpSuccess[PassChangedSuccess].Send(c)
	}
}

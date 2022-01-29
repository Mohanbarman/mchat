package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	services "mchat.com/api/modules/auth/services"
)

type AuthController struct {
	Config  *config.Config
	DB      *gorm.DB
	Service *services.AuthService
}

func (ctrl *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrl.Service.Login(c, &services.JwtService{Config: &ctrl.Config.Jwt})
	}
}

func (ctrl *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrl.Service.Register(c)
	}
}

func (ctrl *AuthController) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrl.Service.GetMe(c)
	}
}

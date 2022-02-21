package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/middlewares"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	jwtService := lib.Jwt{Config: &config.Jwt}

	authCtrl := Controller{
		Service: &Service{
			Config: config,
			Db:     db,
			Jwt:    &jwtService,
			Redis:  lib.GetRedisClient("reset_password", &config.Redis),
			Mail:   &lib.MailClient{Config: &config.SmtpOtp},
		},
	}

	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &jwtService,
		DB:  db,
	}

	router.POST("/login", authCtrl.Login())
	router.POST("/register", authCtrl.Register())
	router.POST("/reset-password/send-mail", authCtrl.SendResetPasswordMail())
	router.GET("/me", authMiddleware.Validate(lib.AccessToken), authCtrl.GetMe())
	router.POST("/reset-password", authCtrl.ResetPassword())
}

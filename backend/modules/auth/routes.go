package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/middlewares"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	jwtService := jwt.JwtService{Config: &config.Jwt}

	authCtrl := Controller{
		Config: config,
		DB:     db,
		Service: &Service{
			Config: config,
			Db:     db,
		},
		OtpSmtp:    &lib.MailClient{Config: &config.SmtpOtp},
		Redis:      lib.GetRedisClient("reset_password", &config.Redis),
		JwtService: &jwtService,
	}

	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &jwtService,
		DB:  db,
	}

	router.POST("/login", authCtrl.Login())
	router.POST("/register", authCtrl.Register())
	router.POST("/reset-password/send-mail", authCtrl.SendResetPasswordMail())
	router.GET("/me", authMiddleware.Validate(jwt.AccessToken), authCtrl.GetMe())
	router.POST("/reset-password", authCtrl.ResetPassword())
}

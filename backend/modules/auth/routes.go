package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/middlewares"
	controllers "mchat.com/api/modules/auth/controllers"
	dto "mchat.com/api/modules/auth/dto"
	services "mchat.com/api/modules/auth/services"
	utils "mchat.com/api/utils"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	authCtrl := controllers.AuthController{
		Config: config,
		DB:     db,
		Service: &services.AuthService{
			Config: config,
			Db:     db,
		},
		OtpSmtp: &utils.MailClient{Config: &config.SmtpOtp},
		Redis:   utils.GetRedisClient("reset_password", &config.Redis),
	}

	jwtService := services.JwtService{Config: &config.Jwt}

	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &jwtService,
		DB:  db,
	}

	router.POST("/login", middlewares.Validate(&dto.LoginDto{}), authCtrl.Login())
	router.POST("/register", middlewares.Validate(&dto.RegisterDto{}), authCtrl.Register())
	router.POST("/reset-password/send-mail", middlewares.Validate(&dto.ResetPasswordDTO{}), authCtrl.SendResetPasswordMail())
	router.GET("/me", authMiddleware.Validate(services.AccessToken), authCtrl.GetMe())
	router.POST("/reset-password", middlewares.Validate(&dto.ResetPasswordChangeDTO{}), authCtrl.ResetPassword())
}

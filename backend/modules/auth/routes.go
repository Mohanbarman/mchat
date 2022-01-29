package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/middlewares"
	controllers "mchat.com/api/modules/auth/controllers"
	dto "mchat.com/api/modules/auth/dto"
	services "mchat.com/api/modules/auth/services"
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
	}

	jwtService := services.JwtService{Config: &config.Jwt}

	router.POST("/login", middlewares.Validate(&dto.LoginDto{}), authCtrl.Login())
	router.POST("/register", middlewares.Validate(&dto.RegisterDto{}), authCtrl.Register())
	router.GET("/me", middlewares.AuthMiddleware(&jwtService, services.AccessToken, db), authCtrl.GetMe())
}

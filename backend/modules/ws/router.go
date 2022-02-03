package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/middlewares"
	auth "mchat.com/api/modules/auth/services"
	"mchat.com/api/modules/ws/connection"
)

func Init(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	ctrl := WsController{
		Upgrader: websocket.Upgrader{},
		DB:       db,
		Config:   config,
	}
	midd := middlewares.AuthMiddleware{
		Jwt: &auth.JwtService{Config: &config.Jwt},
		DB:  db,
	}
	conManager := connection.NewManager()

	router.GET("/connect", midd.Validate(auth.AccessToken), ctrl.CreateConnection(conManager))
}

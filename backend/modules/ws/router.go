package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/middlewares"
	"mchat.com/api/modules/ws/connection"
)

func Init(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB, wsStore *connection.ConnStore) {
	router := rg.Group(prefix)

	ctrl := WsController{
		Upgrader: websocket.Upgrader{},
		DB:       db,
		Config:   config,
	}
	midd := middlewares.AuthMiddleware{
		Jwt: &jwt.JwtService{Config: &config.Jwt},
		DB:  db,
	}

	router.GET("/connect", midd.Validate(jwt.AccessToken), ctrl.CreateConnection(wsStore))
}

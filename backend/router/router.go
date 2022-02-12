package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/auth"
	"mchat.com/api/modules/conversations"
	"mchat.com/api/modules/messages"
	"mchat.com/api/modules/ws"
	"mchat.com/api/modules/ws/connection"
)

func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB, wsStore *connection.ConnStore) {
	rg := engine.Group("/api")

	auth.InitRoutes("/auth", rg, config, db)
	conversations.InitRoutes("/conversations", rg, config, db, wsStore)
	messages.InitRoutes("/messages", rg, config, db, wsStore)
	ws.Init("/ws", rg, config, db, wsStore)
}

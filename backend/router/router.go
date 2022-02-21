package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/modules/auth"
	"mchat.com/api/modules/conversations"
	"mchat.com/api/modules/messages"
	"mchat.com/api/ws"
)

func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB, wsStore *lib.WsStore) {
	rg := engine.Group("/api")
	wsrg := engine.Group("/ws")

	auth.InitRoutes("/auth", rg, config, db)
	conversations.InitRoutes("/conversations", rg, config, db, wsStore)
	messages.InitRoutes("/messages", rg, config, db, wsStore)
	ws.Init(wsrg, config, db, wsStore)
}

package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/auth"
	"mchat.com/api/modules/ws"
)

func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB) {
	rg := engine.Group("/api")

	auth.InitRoutes("/auth", rg, config, db)
	ws.Init("/ws", rg, config, db)
}

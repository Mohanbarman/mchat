package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/auth"
)

func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB) {
	rg := engine.Group("/api")

	auth.InitRoutes("/auth", rg, config, db)
}

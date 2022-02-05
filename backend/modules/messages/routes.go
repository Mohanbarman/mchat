package messages

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/ws/connection"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB, wsStore *connection.ConnStore) {
	router := rg.Group(prefix)

	router.GET("")
}

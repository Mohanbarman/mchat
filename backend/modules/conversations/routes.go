package conversations

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/middlewares"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB, wsStore *lib.WsStore) {
	router := rg.Group(prefix)

	ctrl := Controller{
		DB:      db,
		Config:  config,
		Service: &Service{DB: db},
	}
	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &lib.Jwt{Config: &config.Jwt},
		DB:  db,
	}

	router.GET("", authMiddleware.Validate(), ctrl.GetAll())
	router.GET("/:conversation_id", authMiddleware.Validate(), ctrl.GetOne())
	router.POST("", authMiddleware.Validate(), ctrl.Create())
}

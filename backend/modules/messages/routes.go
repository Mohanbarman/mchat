package messages

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/middlewares"
	"mchat.com/api/modules/ws/connection"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB, wsStore *connection.ConnStore) {
	router := rg.Group(prefix)

	ctrl := Controller{DB: db, Config: config}

	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &jwt.JwtService{Config: &config.Jwt},
		DB:  db,
	}

	router.GET("/:conversation_id", authMiddleware.Validate(jwt.AccessToken), ctrl.GetAll())
}

package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
)

func Init(rg *gin.RouterGroup, config *config.Config, db *gorm.DB, wsStore *lib.WsStore) {
	router := rg.Group("")

	ctrl := WsController{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		DB:     db,
		Config: config,
	}

	router.GET("", ctrl.CreateConnection(wsStore))
}

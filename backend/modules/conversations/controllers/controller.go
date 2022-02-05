package conversation

import (
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/modules/ws/connection"
)

type ConversationCtrl struct {
	Config *config.Config
	DB     *gorm.DB
	WS     *connection.ConnStore
}

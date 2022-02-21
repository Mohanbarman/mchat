package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// message status
const (
	MessageStatusSent      int = 0
	MessageStatusDelivered int = 1
	MessageStatusSeen      int = 2
)

type MessageModel struct {
	Base
	UUID           string `gorm:"unique"`
	Text           string `gorm:"type:TEXT"`
	FileID         sql.NullInt64
	ConversationID uint
	File           FileModel
	Conversation   ConversationModel
	FromUserID     int
	ToUserID       int
	FromUser       UserModel
	ToUser         UserModel
	Status         int
}

func (model *MessageModel) BeforeCreate(scope *gorm.DB) (err error) {
	model.UUID = uuid.NewV4().String()
	return
}

func (model *MessageModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"id":         model.ID,
		"text":       model.Text,
		"file_id":    model.FileID.Int64,
		"created_at": model.CreatedAt,
		"status":     model.Status,
	}
}

func (model *MessageModel) TableName() string {
	return "messages"
}

//
// Scopes
//

// update status to seen of messages of user conversation
func SeenConversationMessageStatus(conversationID uint, userCol string, userID uint) gormScope {
	return func(d *gorm.DB) *gorm.DB {
		d.Where("conversation_id = ? AND ?=?", conversationID, userCol, userID)
		return d
	}
}

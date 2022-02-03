package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const (
	MessageStatusSent      int = 0
	MessageStatusDelivered int = 1
	MessageStatusReaded    int = 2
)

type MessageModel struct {
	gorm.Model
	UUID           string `gorm:"unique"`
	Text           string `gorm:"type:TEXT"`
	FileID         sql.NullInt64
	ConversationID uint
	File           FileModel
	Conversation   ConversationModel
	Status         int
}

func (model *MessageModel) BeforeCreate(scope *gorm.DB) (err error) {
	model.UUID = uuid.NewV4().String()
	return
}

func (model *MessageModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"text":         model.Text,
		"file_id":      model.FileID.Int64,
		"conversation": model.Conversation.UUID,
	}
}

func (model *MessageModel) TableName() string {
	return "messages"
}

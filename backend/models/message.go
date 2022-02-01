package models

import "gorm.io/gorm"

type MessageModel struct {
	gorm.Model
	Text           string `gorm:"type:TEXT"`
	FileID         uint
	ConversationID uint
	File           FileModel
	Conversation   ConversationModel
}

func (model *MessageModel) TableName() string {
	return "messages"
}

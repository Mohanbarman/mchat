package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ConversationModel struct {
	gorm.Model
	UUID       string `gorm:"unique;not null"`
	FromUserID uint
	FromUser   UserModel
	ToUserID   uint
	ToUser     UserModel
}

func (model *ConversationModel) TableName() string {
	return "conversations"
}

func (model *ConversationModel) BeforeCreate(scope *gorm.DB) (err error) {
	model.UUID = uuid.NewV4().String()
	return
}

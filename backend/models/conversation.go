package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ConversationModel struct {
	Base
	UUID                string `gorm:"unique;not null"`
	FromUserID          uint
	FromUser            UserModel
	ToUserID            uint
	ToUser              UserModel
	FromUserUnreadCount uint
	ToUserUnreadCount   uint
	LastMessageText     string
	LastMessageTime     sql.NullTime
}

func (model *ConversationModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"id":              model.UUID,
		"created_at":      model.CreatedAt,
		"updated_at":      model.UpdatedAt,
		"lastMessage":     model.LastMessageText,
		"lastMessageTime": model.LastMessageTime.Time,
	}
}

func (model *ConversationModel) TableName() string {
	return "conversations"
}

func (model *ConversationModel) BeforeCreate(scope *gorm.DB) (err error) {
	model.UUID = uuid.NewV4().String()
	return
}

//
// Scopes
//

// Get one to one user conversation
func FindUserConversation(user1 uint, user2 uint) gormScope {
	return func(d *gorm.DB) *gorm.DB {
		d.Where(
			"(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
			user1, user2, user2, user1,
		)
		return d
	}
}

// Get all conversations of the user
func FindAllUserConversations(id uint) gormScope {
	return func(d *gorm.DB) *gorm.DB {
		d.Where("from_user_id = ? OR to_user_id = ?", id, id).Preload(clause.Associations).Order("created_at desc")
		return d
	}
}

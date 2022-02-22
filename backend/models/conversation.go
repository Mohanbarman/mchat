package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ConversationModel struct {
	BaseModel
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
	data := map[string]interface{}{
		"id":                model.UUID,
		"created_at":        model.CreatedAt,
		"updated_at":        model.UpdatedAt,
		"last_message":      model.LastMessageText,
		"last_message_time": model.LastMessageTime.Time,
	}

	return data
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
		d.Where("from_user_id = ? OR to_user_id = ?", id, id).Preload(clause.Associations).Order("updated_at desc")
		return d
	}
}

//
// helpers
//

func (model *ConversationModel) GetOtherUser(forUser uint) uint {
	otherUser := model.FromUserID

	if model.FromUserID == forUser {
		otherUser = model.ToUserID
	}

	return otherUser
}

func (model *ConversationModel) TransformForUser(forUserID uint) map[string]interface{} {
	data := map[string]interface{}{
		"id":                model.UUID,
		"created_at":        model.CreatedAt,
		"updated_at":        model.UpdatedAt,
		"last_message":      model.LastMessageText,
		"last_message_time": model.LastMessageTime.Time,
	}

	if model.ToUserID == forUserID {
		data["user"] = model.ToUser.Transform()
		data["is_unread"] = model.ToUserUnreadCount > 0
		data["unread_count"] = model.ToUserUnreadCount
	} else {
		data["user"] = model.FromUser.Transform()
		data["is_unread"] = model.FromUserUnreadCount > 0
		data["unread_count"] = model.FromUserUnreadCount
	}

	return data
}

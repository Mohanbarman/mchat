package users

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UUID           string `gorm:"column:unique"`
	Email          string `gorm:"column:email;unique"`
	Password       string
	Name           string
	Status         string
	ProfilePicture string
}

func (user *UserModel) BeforeCreate(scope *gorm.DB) (err error) {
	user.UUID = uuid.NewV4().String()
	return
}

func (user *UserModel) TableName() string {
	return "users"
}

func (user *UserModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"id":         user.UUID,
		"email":      user.Email,
		"name":       user.Name,
		"profile":    user.ProfilePicture,
		"status":     user.Status,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
}

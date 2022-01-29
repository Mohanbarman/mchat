package db

import (
	"gorm.io/gorm"
	user_models "mchat.com/api/modules/users/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&user_models.UserModel{})
}

package db

import (
	"gorm.io/gorm"
	"mchat.com/api/models"
)

func RunMigrations(db *gorm.DB) {
	m := []interface{}{
		&models.UserModel{},
		&models.ConversationModel{},
		&models.FileModel{},
		&models.MessageModel{},
		&models.UserModel{},
	}
	db.AutoMigrate(m...)
}

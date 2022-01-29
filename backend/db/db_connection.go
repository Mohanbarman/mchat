package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mchat.com/api/config"
)

func New(config config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}

	RunMigrations(db)

	return db
}

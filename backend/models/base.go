package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type gormScope = func(*gorm.DB) *gorm.DB

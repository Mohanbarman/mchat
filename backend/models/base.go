package models

import (
	"database/sql"
	"time"
)

type Base struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type FileModel struct {
	gorm.Model
	MimeType  string
	FilePath  string
	SizeBytes int64
	FileHash  string
	UserID    sql.NullInt64
}

func (model *FileModel) TableName() string {
	return "files"
}

package models

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	MimeType  string
	FilePath  string
	SizeBytes int64
	FileHash  string
}

func (model *FileModel) TableName() string {
	return "files"
}

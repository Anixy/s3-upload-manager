package domain

import (
	"github.com/Anixy/s3-upload-manager/model"

	"gorm.io/gorm"
)

type Database interface {
	AddFile(file *model.File) (err error)
	DeleteFileByID(fileID uint) (err error)
}

type databaseImpl struct {
	*gorm.DB
}

func NewDatabase(db *gorm.DB) Database {
	return &databaseImpl{
		DB: db,
	}
}

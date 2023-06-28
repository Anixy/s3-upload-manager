package service

import (
	"os"

	"github.com/Anixy/s3-upload-manager/repository/domain"

	"github.com/Anixy/s3-upload-manager/config"
)

type File interface {
	UploadFile(file *os.File, size int64,contentType string) (err error)
}

type fileImpl struct {
	config.Config
	DatabaseDomain domain.Database
	S3Domain domain.S3
}

func NewFile(config config.Config, databaseDomain domain.Database, s3Domain domain.S3) File {
	return &fileImpl{
		Config: config,
		DatabaseDomain: databaseDomain,
		S3Domain: s3Domain,
	}
}



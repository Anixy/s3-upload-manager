package service

import (
	"os"
	"testing"

	"github.com/Anixy/s3-upload-manager/repository/domain"

	"github.com/Anixy/s3-upload-manager/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_File_UploadFile_Success(t *testing.T) {
	cfg := config.Config{
		BucketName: "test-bucket",
	}
	s3Domain := new(domain.MockS3)
	databaseDomain := new(domain.MockDatabase)
	filePath := "data/file.txt"
	mockFile, err := os.Open(filePath)
	assert.NoError(t, err)
	mockStat, err := mockFile.Stat()
	assert.NoError(t, err)
	contentType := "text/plain"
	fileService := NewFile(cfg, databaseDomain, s3Domain)
	s3Domain.On("UploadObject", mock.Anything, mockFile, mockStat.Size(), mock.Anything).Return(mockStat.Size(), nil)
	databaseDomain.On("AddFile", mock.Anything).Return(nil)
	err = fileService.UploadFile(mockFile, mockStat.Size(), contentType)
	assert.NoError(t, err)
}
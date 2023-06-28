package domain

import (
	"errors"
	"os"
	"testing"

	"github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
)

func Test_S3_UploadObject_Success(t *testing.T) {
	mockS3Domain := new(MockS3)

	filePath := "mock/data/file.txt"
	fileName := "file.txt"

	mockFile, err := os.Open(filePath)
	assert.NoError(t, err)

	mockStat, err := mockFile.Stat()
	assert.NoError(t, err)

	opts := minio.PutObjectOptions{
		ContentType: "text/plain",
	}
	mockS3Domain.On("UploadObject", fileName, mockFile, mockStat.Size(), opts).Return(mockStat.Size(), nil)
	n, err := mockS3Domain.UploadObject(fileName, mockFile, mockStat.Size(), opts)
	assert.NoError(t, err)
	assert.Equal(t, mockStat.Size(), n)
	mockS3Domain.AssertExpectations(t)
}

func Test_S3_UploadObject_Failed(t *testing.T) {
	mockS3Domain := new(MockS3)
	// bucketName := "test_bucket"
	filePath := "mock/data/file.txt"
	fileName := "file.txt"

	mockFile, err := os.Open(filePath)
	assert.NoError(t, err)

	mockStat, err := mockFile.Stat()
	assert.NoError(t, err)

	opts := minio.PutObjectOptions{
		ContentType: "text/plain",
	}

	errMsg := errors.New("Failed upload object")
	mockS3Domain.On("UploadObject", fileName, mockFile, mockStat.Size(), opts).Return(int64(0), errMsg)
	n, err := mockS3Domain.UploadObject(fileName, mockFile, mockStat.Size(), opts)
	assert.Error(t, err)
	assert.NotEqual(t, mockStat.Size(), n)
	mockS3Domain.AssertExpectations(t)
}

func Test_S3_DeleteObject_Success(t *testing.T) {
	mockS3Domain := new(MockS3)
	fileName := "file.txt"
	mockS3Domain.On("DeleteObject", fileName).Return(nil)
	err := mockS3Domain.DeleteObject(fileName)
	assert.NoError(t, err)
	mockS3Domain.AssertExpectations(t)
}

func Test_S3_DeleteObject_Failed(t *testing.T) {
	mockS3Domain := new(MockS3)
	fileName := "file.txt"
	errMsg := errors.New("Failed delete object")
	mockS3Domain.On("DeleteObject", fileName).Return(errMsg)
	err := mockS3Domain.DeleteObject(fileName)
	assert.Error(t, err)
	mockS3Domain.AssertExpectations(t)
}
package domain

import (
	"errors"
	"testing"

	"github.com/Anixy/s3-upload-manager/model"

	"github.com/stretchr/testify/assert"
)

func Test_Database_AddFile_Success(t *testing.T) {
	mockDatabaseDomain := new(MockDatabase)
	file := model.File{
		Name:       "test.txt",
		Visibility: "public",
		Size:       40,
		Bucket:     "test-bucket",
	}
	mockDatabaseDomain.On("AddFile", &file).Return(nil)
	err := mockDatabaseDomain.AddFile(&file)
	assert.NoError(t, err)
	mockDatabaseDomain.AssertExpectations(t)
}

func Test_Database_AddFile_Failed(t *testing.T) {
	mockDatabaseDomain := new(MockDatabase)
	file := model.File{
		Name:       "test.txt",
		Visibility: "public",
		Size:       40,
		Bucket:     "test-bucket",
	}
	errMsg := errors.New("Failed add file record")
	mockDatabaseDomain.On("AddFile", &file).Return(errMsg)
	err := mockDatabaseDomain.AddFile(&file)
	assert.Error(t, err)
	mockDatabaseDomain.AssertExpectations(t)
}

func Test_Database_DeleteFileByID_Success(t *testing.T) {
	mockDatabaseDomain := new(MockDatabase)
	mockDatabaseDomain.On("DeleteFileByID", uint(1)).Return(nil)
	err := mockDatabaseDomain.DeleteFileByID(1)
	assert.NoError(t, err)
	mockDatabaseDomain.AssertExpectations(t)
}

func Test_Database_DeleteFileByID_Failed(t *testing.T) {
	mockDatabaseDomain := new(MockDatabase)
	errMsg := errors.New("Failed delete file record")
	mockDatabaseDomain.On("DeleteFileByID", uint(1)).Return(errMsg)
	err := mockDatabaseDomain.DeleteFileByID(1)
	assert.Error(t, err)
	mockDatabaseDomain.AssertExpectations(t)
}

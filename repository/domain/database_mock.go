package domain

import (
	"time"

	"github.com/Anixy/s3-upload-manager/model"

	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) AddFile(file *model.File) (err error) {
	args := m.Called(file)
	if file.ID == 0 {
		file.ID = 1
	}
	file.CreatedAt = time.Now()
	file.UpdatedAt = time.Now()
	return args.Error(0)
}

func (m *MockDatabase) DeleteFileByID(fileID uint) (err error) {
	args := m.Called(fileID)
	return args.Error(0)
}

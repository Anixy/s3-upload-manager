package domain

import (
	"io"

	"github.com/minio/minio-go"
	"github.com/stretchr/testify/mock"
)

type MockS3 struct {
	mock.Mock
}

// PutObject is a mock implementation of the PutObject method of the Minio client
func (m *MockS3) UploadObject(objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error) {
	args := m.Called(objectName, reader, objectSize, opts)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockS3) DeleteObject(objectName string) (err error) {
	args := m.Called(objectName)
	return args.Error(0)
}
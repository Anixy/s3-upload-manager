package domain

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMinioClient is a mock implementation of the Minio client
type MockMinioClient struct {
	mock.Mock
}

// PutObject is a mock implementation of the PutObject method of the Minio client
func (m *MockMinioClient) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error) {
	args := m.Called(bucketName, objectName, reader, objectSize, opts)
	return args.Get(0).(int64), args.Error(1)
}

// TestS3Domain_UploadFile tests the UploadObject method of the S3Domain struct
func TestS3Domain_UploadFile(t *testing.T) {
	mockClient := new(MockMinioClient)
	bucketName := "test-bucket"
	objectName := "file.txt"
	filePath := "mock/data/file.txt"
	// Open the mock file and get its stats
	mockFile, err := os.Open(filePath)
	assert.NoError(t, err)
	mockStat, err := mockFile.Stat()
	assert.NoError(t, err)
	// Set the options for the PutObject call
	putObjectOptions := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}
	// Set up the mock client to expect a PutObject call with the specified arguments
	mockClient.On("PutObject", bucketName, objectName, mock.Anything, mockStat.Size(), putObjectOptions).Return(mockStat.Size(), nil)
	// Create a new S3Domain instance and call the UploadObject method
	s3Domain := NewS3(mockClient, bucketName)
	n, err := s3Domain.UploadObject(objectName, mockFile, mockStat.Size(), putObjectOptions)
	// Assert that no error occurred and that the expected number of bytes were uploaded
	assert.NoError(t, err)
	assert.Equal(t, mockStat.Size(), n)
	// Assert that the mock client received the expected calls
	mockClient.AssertExpectations(t)
}

// TestS3Domain_UploadFile_Error tests the UploadObject method of the S3Domain struct when an error occurs
func TestS3Domain_UploadFile_Error(t *testing.T) {
	mockClient := new(MockMinioClient)
	bucketName := "test-bucket"
	objectName := "file.txt"
	filePath := "mock/data/file.txt"
	// Open the mock file and get its stats
	mockFile, err := os.Open(filePath)
	assert.NoError(t, err)
	mockStat, err := mockFile.Stat()
	assert.NoError(t, err)
	// Set the options for the PutObject call
	putObjectOptions := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}
	// Set up the mock client to expect a PutObject call with the specified arguments and return an error
	expectedError := errors.New("upload failed")
	mockClient.On("PutObject", bucketName, objectName, mock.Anything, mockStat.Size(), putObjectOptions).Return(int64(0), expectedError)
	// Create a new S3Domain instance and call the UploadObject method
	s3Domain := NewS3(mockClient, bucketName)
	_, err = s3Domain.UploadObject(objectName, mockFile, mockStat.Size(), putObjectOptions)
	// Assert that the error returned matches the expected error and that the mock client received the expected calls
	assert.EqualError(t, err, expectedError.Error())
	mockClient.AssertExpectations(t)
}

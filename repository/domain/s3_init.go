// Package domain provides an interface for interacting with Amazon S3
package domain

import (
	"io"

	"github.com/minio/minio-go"
)

// S3 is an interface for interacting with Amazon S3
type S3 interface {
	UploadObject(objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error)
}

// s3Impl is a concrete implementation of the S3 interface
type s3Impl struct {
	minio.Client
	BucketName string
}

// NewS3 initializes a new S3 implementation and returns it as an S3 interface
func NewS3(minioClient minio.Client, bucketName string) S3 {
	return &s3Impl{
		Client:     minioClient,
		BucketName: bucketName,
	}
}

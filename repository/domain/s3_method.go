package domain

import (
	"io"
	"log"

	"github.com/minio/minio-go"
)

// UploadObject uploads an object to the specified S3 bucket with the given object name, reader, size and options.
func (d *s3Impl) UploadObject(objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error) {
	n, err = d.Client.PutObject(d.BucketName, objectName, reader, objectSize, opts)
	if err != nil {
		log.Println(err)
	}
	return
}
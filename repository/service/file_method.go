package service

import (
	"os"

	"github.com/Anixy/s3-upload-manager/util"

	"github.com/Anixy/s3-upload-manager/model"

	"github.com/minio/minio-go"
)

func (s *fileImpl) UploadFile(file *os.File, size int64,contentType string) (err error) {
	objectName, err  := util.GenerateRandomString(20)
	if err != nil {
		return
	}
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		UserMetadata: map[string]string{
			"X-Amz-Acl": "public-read",
		},
	}
	n, err := s.S3Domain.UploadObject(objectName, file, size, opts)
	if err != nil {
		return
	}
	fileDB := model.File{
		Name: objectName,
		Visibility: "public",
		Size: n,
		Bucket: s.Config.BucketName,
	}
	err = s.DatabaseDomain.AddFile(&fileDB)
	if err != nil {
		return
	}
	return
}

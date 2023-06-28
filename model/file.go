package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string
	Visibility string
	Size int64
	Bucket string
	Url string
}
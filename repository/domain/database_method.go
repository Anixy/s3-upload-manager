package domain

import "github.com/Anixy/s3-upload-manager/model"

func (d *databaseImpl) AddFile(file *model.File) (err error) {
	err = d.DB.Create(&file).Error
	if err != nil {
		return
	}
	return
}

func (d *databaseImpl) DeleteFileByID(fileID uint) (err error) {
	err = d.DB.Delete(&model.File{}, fileID).Error
	if err != nil {
		return
	}
	return
}
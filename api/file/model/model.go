package model

import (
	"uploader-service/config/tools"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var model *gorm.DB
var table *gorm.DB

func SetUpModel() {
	model = tools.DB
	table = tools.DB.Table("files")
	model.AutoMigrate(&File{})
}

func CreateFile(fileNewDTO FileNewDTO) (File, error) {
	newFile := FromNewToFile(fileNewDTO)

	if result := model.Create(&newFile); result.Error != nil {
		return File{}, result.Error
	}

	return newFile, nil
}

func GetFilesByUserId(userId string) ([]File, error) {
	var files []File

	userIdUUID := uuid.FromStringOrNil(userId)

	if err := model.Where(&File{UserID: userIdUUID}).Find(&files).Error; err != nil {
		return files, err
	}

	return files, nil
}

func GetFileByUserId(fileId string, userId string) (File, error) {
	fileIdUUID := uuid.FromStringOrNil(fileId)
	userIdUUID := uuid.FromStringOrNil(userId)

	file := File{
		ID:     fileIdUUID,
		UserID: userIdUUID,
	}

	if err := model.First(&file).Error; err != nil {
		return file, err
	}

	return file, nil
}

func UpdateFile() {

}

func DeleteFile(fileId string) (File, error) {
	fileIdUUID := uuid.FromStringOrNil(fileId)
	fileToRemove := File{ID: fileIdUUID}

	if err := model.Delete(&fileToRemove).Error; err != nil {
		return fileToRemove, err
	}

	return fileToRemove, nil
}

package model

import (
	"time"
	"uploader-service/config/tools"

	"gorm.io/gorm"
)

type File struct {
	ID        uint      // Standard field for the primary key
	Name      string    // A regular string field
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

var model *gorm.DB

func SetUpModel() {
	model = tools.DB
}

func CreateFile() {

}

func GetFilesByUser() {

}

func GetFile() {

}

func UpdateFile() {

}

func DeleteFile() {

}

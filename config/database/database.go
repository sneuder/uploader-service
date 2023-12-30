package database

import (
	userModel "uploader-service/api/user/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("uploader.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func SetModels() {
	userModel.SetModel()
}

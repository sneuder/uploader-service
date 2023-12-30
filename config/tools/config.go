package tools

import (
	"uploader-service/services/fileService"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var E *echo.Echo
var G *echo.Group
var DB *gorm.DB

func SetGlobalTools(echo *echo.Echo, group *echo.Group, db *gorm.DB) {
	E = echo
	G = group
	DB = db
}

func SetBaseFolder() error {
	if err := fileService.CreateFolder("", fileService.BaseDirectory); err != nil {
		return err
	}

	return nil
}

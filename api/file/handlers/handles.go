package handlers

import (
	"net/http"
	"uploader-service/api/file/service"

	"github.com/labstack/echo/v4"
)

func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to get the file from the request")
	}

	message, err := service.SaveFile(file)

	return c.String(http.StatusOK, message)
}

func DownloadFile(c echo.Context) error {
	return c.JSON(http.StatusOK, "asa")
}

func RenameFile(c echo.Context) error {
	return c.JSON(http.StatusOK, "asa")
}

func DeleteFile(c echo.Context) error {
	return c.JSON(http.StatusOK, "asa")
}

func DeleteFiles(c echo.Context) error {
	return c.JSON(http.StatusOK, "asa")
}

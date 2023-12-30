package handlers

import (
	"errors"
	"net/http"
	"uploader-service/api/file/model"
	"uploader-service/crash"
	"uploader-service/services/fileService"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToGetFile, err))
	}

	nameFile := c.FormValue("name")
	userId := c.FormValue("userId")

	fileNewDTO := model.FileNewDTO{
		Name:   nameFile,
		UserID: uuid.FromStringOrNil(userId),
	}

	createdFile, err := model.CreateFile(fileNewDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToProcessFile, err))
	}

	err = fileService.SaveFile(file, createdFile.UserID.String(), createdFile.ID.String())
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToProcessFile, err))
	}

	return c.JSON(http.StatusOK, createdFile)
}

func GetAllFilesByUserId(c echo.Context) error {
	userId := c.Param("userId")

	files, err := model.GetFilesByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToProcessFile, err))
	}

	return c.JSON(http.StatusOK, files)
}

func GetFilesByUserId(c echo.Context) error {
	fileId := c.Param("fileId")
	userId := c.Param("userId")

	file, err := model.GetFileByUserId(fileId, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToGetFile, err))
	}

	return c.JSON(http.StatusOK, file)
}

func GetFile(c echo.Context) error {
	fileId := c.Param("fileId")
	userId := c.Param("userId")

	filePath, err := fileService.FindFilePath(userId, fileId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToGetFile, err))
	}

	if filePath == "" {
		err := errors.New("Could not find file")
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToGetFile, err))
	}

	return c.File(filePath)
}

// func UpdateFile(c echo.Context) error {

// }

func DeleteFile(c echo.Context) error {
	fileId := c.Param("fileId")
	userId := c.Param("userId")

	fileRemoved, err := model.DeleteFile(fileId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToRemoveFile, err))
	}
	if err := fileService.RemoveFile(userId, fileId); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToRemoveFile, err))
	}

	return c.JSON(http.StatusOK, fileRemoved)
}

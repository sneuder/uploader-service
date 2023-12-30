package fileService

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var BaseDirectory string = "uploads"

func SaveFile(file *multipart.FileHeader, userId string) error {
	if err := CreateFolder(BaseDirectory, userId); err != nil {
		return err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a destination file on the server directly in the "uploads" folder
	dstPath := filepath.Join("uploads", userId, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded file directly to the destination folder
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func CreateFolder(baseDirectory string, folderName string) error {
	directoryPath := filepath.Join(baseDirectory, folderName)

	_, err := os.Stat(directoryPath)
	if errDir := os.MkdirAll(directoryPath, 0755); errDir != nil {
		return err
	}

	return nil
}

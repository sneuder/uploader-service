package fileService

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var BaseDirectory string = "uploads"

func SaveFile(file *multipart.FileHeader, userId string, fileId string) error {
	if err := CreateFolder(BaseDirectory, userId); err != nil {
		return err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileExtension := filepath.Ext(file.Filename)
	filename := fileId + fileExtension

	// Create a destination file on the server directly in the "uploads" folder
	dstPath := filepath.Join(BaseDirectory, userId, filename)
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

func FindFilePath(folderName string, fileName string) (string, error) {
	files, err := filepath.Glob(filepath.Join(BaseDirectory, folderName, fileName+".*"))
	if err != nil || len(files) == 0 {
		return "", err
	}
	return files[0], nil
}

func RemoveFile(folderName string, fileName string) error {
	filePath, err := FindFilePath(folderName, fileName)

	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

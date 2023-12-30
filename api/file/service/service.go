package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(file *multipart.FileHeader) (string, error) {

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "Failed to open the uploaded file", err
	}
	defer src.Close()

	// Create a destination file on the server directly in the "uploads" folder
	dstPath := filepath.Join("uploads", file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "Failed to create the destination file", err
	}
	defer dst.Close()

	// Copy the uploaded file directly to the destination folder
	if _, err = io.Copy(dst, src); err != nil {
		return "Failed to copy the file to the destination", err
	}

	return "File uploaded", nil
}

func createFolder() {

}

func folderExists() {

}

func RenameFile() {

}

func DeleteFile() {

}

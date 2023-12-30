package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type File struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FileNewDTO struct {
	UserID uuid.UUID `json:"userId" validate:"required"`
	Name   string    `json:"name" validate:"required"`
}

// dto mapping

func FromNewToFile(fileNewDTO FileNewDTO) File {
	file := File{
		ID:     uuid.NewV4(),
		UserID: fileNewDTO.UserID,
		Name:   fileNewDTO.Name,
	}

	return file
}

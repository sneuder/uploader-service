package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserSignUpDTO struct {
	Name            string `json:"name" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}

type UserSignInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserUpdateDTO struct {
	ID       uuid.UUID `param:"id"`
	Name     *string   `json:"name,omitempty"`
	LastName *string   `json:"lastName,omitempty"`
}

// mapping dtos

func FromSignUpToUser(userSignUpDTO *UserSignUpDTO) User {
	newUser := User{
		ID:       uuid.NewV4(),
		Name:     userSignUpDTO.Name,
		LastName: userSignUpDTO.LastName,
		Email:    userSignUpDTO.Email,
		Password: userSignUpDTO.Password,
	}

	return newUser
}

func FromUserToResponse(user *User) UserResponseDTO {
	userReponseDTO := UserResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	return userReponseDTO
}

func FromSignInToUser(userSignInDTO *UserSignInDTO) User {
	user := User{
		Email:    userSignInDTO.Email,
		Password: userSignInDTO.Password,
	}

	return user
}

func FromUserUpdateToUser() {

}

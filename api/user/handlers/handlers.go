package handlers

import (
	"errors"
	"net/http"
	"uploader-service/api/user/model"
	"uploader-service/api/user/services/crypt"
	"uploader-service/crash"
	"uploader-service/services/authService"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	userDTO := new(model.UserSignUpDTO)

	if err := c.Bind(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.MissingSomeFields, err))
	}

	if err := c.Validate(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.MissingSomeFields, err))
	}

	if userDTO.Password != userDTO.PasswordConfirm {
		err := errors.New("Different passwords in user creating")
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.DifferentPassword, err))
	}

	hashedPassword, _ := crypt.HashPassword(userDTO.Password)
	userDTO.Password = hashedPassword

	createdUser, err := model.SignUp(userDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.DifferentFieldTypes, err))
	}

	userResponseDT := model.FromUserToResponse(&createdUser)
	return c.JSON(http.StatusOK, userResponseDT)
}

func SignIn(c echo.Context) error {
	userDTO := new(model.UserSignInDTO)

	if err := c.Bind(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.MissingSomeFields, err))
	}

	if err := c.Validate(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.MissingSomeFields, err))
	}

	hashedPassword, _ := crypt.HashPassword(userDTO.Password)
	userDTO.Password = hashedPassword

	foundUser, err := model.GetUserByEmailPassword(userDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.NotFoundUserInAuth, err))
	}

	token, err := authService.GenerateToken(foundUser.ID)
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+token)

	userResponseDTO := model.FromUserToResponse(&foundUser)
	return c.JSON(http.StatusOK, userResponseDTO)
}

func UpdateUser(c echo.Context) error {
	updateUser := new(model.UserUpdateDTO)

	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.MissingSomeFields, err))
	}

	updatedUser, err := model.UpdateUser(*updateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, crash.GenerateError(crash.FailToUpadeUSer, err))
	}

	userResponseDTO := model.FromUserToResponse(&updatedUser)
	return c.JSON(http.StatusOK, userResponseDTO)
}

func RemoveUser(c echo.Context) {

}

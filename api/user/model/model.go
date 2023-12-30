package model

import (
	"uploader-service/config/tools"

	"gorm.io/gorm"
)

var model *gorm.DB
var table *gorm.DB

func SetModel() {
	model = tools.DB
	table = tools.DB.Table("users")
	model.AutoMigrate(&User{})
}

func SignUp(userSignUpDTO *UserSignUpDTO) (User, error) {
	newUser := FromSignUpToUser(userSignUpDTO)

	if result := model.Create(&newUser); result.Error != nil {
		return User{}, result.Error
	}

	return newUser, nil
}

func GetUserByEmailPassword(userSignInDTO *UserSignInDTO) (User, error) {
	foundUser := FromSignInToUser(userSignInDTO)

	result := model.First(&foundUser)
	if result.Error != nil {
		return User{}, result.Error
	}

	return foundUser, nil
}

func UpdateUser(userUpdateDTO UserUpdateDTO) (User, error) {
	foundUser := User{}
	foundUser.ID = userUpdateDTO.ID

	if err := model.First(&foundUser).Error; err != nil {
		return foundUser, err
	}

	if userUpdateDTO.Name != nil {
		foundUser.Name = *userUpdateDTO.Name
	}

	if userUpdateDTO.LastName != nil {
		foundUser.LastName = *userUpdateDTO.LastName
	}

	if err := model.Save(&foundUser).Error; err != nil {
		return foundUser, err
	}

	return foundUser, nil
}

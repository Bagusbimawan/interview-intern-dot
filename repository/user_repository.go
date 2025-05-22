package repository

import (
	"intervew-intern-dot/config"
	"intervew-intern-dot/model"
)

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(ID int) (*model.User, error) {
	var user model.User
	err := config.DB.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

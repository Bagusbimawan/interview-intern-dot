package repository

import (
	"errors"
	"intervew-intern-dot/config"
	"intervew-intern-dot/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ValidateUser(email, password string) (bool, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func IsEmailExists(email string) (bool, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return user.ID != 0, nil
}

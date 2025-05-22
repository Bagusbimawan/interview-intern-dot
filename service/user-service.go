package service

import (
	"errors"
	"intervew-intern-dot/model"
	"intervew-intern-dot/repository"
	"intervew-intern-dot/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *model.User) error {
	exist, err := repository.IsEmailExists(user.Email)

	if err != nil {
		return err
	}

	if exist {
		return errors.New("email already exists")
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	return repository.CreateUser(user)
}

func LoginUser(user *model.User) (*model.User, string, error) {
	valid, err := repository.ValidateUser(user.Email, user.Password)
	if err != nil {
		return nil, "", errors.New("email not found")
	}

	if !valid {
		return nil, "", errors.New("invalid Password")
	}

	dbUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	token, err := utils.GenerateJWT(int(dbUser.ID), dbUser.Email)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	dbUser.Password = ""
	return dbUser, token, nil
}

func GetUserProfile(id int) (*model.User, error) {
	return repository.GetUserByID(id)
}

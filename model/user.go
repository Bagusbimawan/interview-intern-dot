package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Projects []Project `json:"projects" gorm:"foreignKey:UserID"`
}

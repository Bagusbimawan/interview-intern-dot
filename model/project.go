package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
}

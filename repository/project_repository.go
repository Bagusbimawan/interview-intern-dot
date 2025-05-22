package repository

import (
	"intervew-intern-dot/config"
	"intervew-intern-dot/model"
)

func CreateProject(project *model.Project) error {
	return config.DB.Create(project).Error
}

func GetProjectsByUserID(userID int) ([]model.Project, error) {
	var projects []model.Project
	err := config.DB.Preload("Tasks").Where("user_id = ?", userID).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectDetailByID(id int) (*model.Project, error) {
	var project model.Project
	err := config.DB.Preload("Tasks").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func UpdateProject(project *model.Project) error {
	return config.DB.Save(project).Error
}

func UpdateByID(id int, updatedProject *model.Project) error {
	var project model.Project
	// Temukan proyek berdasarkan ID
	if err := config.DB.First(&project, id).Error; err != nil {
		return err
	}
	// Update field yang diizinkan
	project.Title = updatedProject.Title
	project.Description = updatedProject.Description
	// Tambahkan field lain jika perlu

	return config.DB.Save(&project).Error
}

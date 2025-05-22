package repository

import (
	"intervew-intern-dot/config"
	"intervew-intern-dot/model"
)

func CreateTask(task *model.Task) error {
	return config.DB.Create(task).Error
}

func GetTasksByProjectID(projectID int) ([]model.Task, error) {
	var tasks []model.Task
	err := config.DB.Where("project_id = ?", projectID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

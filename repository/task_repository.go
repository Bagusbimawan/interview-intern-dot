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

func UpdateTaskByID(id int, updatedTask *model.Task) error {
	var task model.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.IsDone = updatedTask.IsDone
	return config.DB.Save(&task).Error
}

func DeleteTaskByID(id int) error {
	var task model.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&task).Error
}

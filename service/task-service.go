package service

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/repository"
)

func CreateTask(task *model.Task) error {
	return repository.CreateTask(task)
}

func UpdateTaskByID(id int, updatedTask *model.Task) error {
	return repository.UpdateTaskByID(id, updatedTask)
}

func DeleteTaskByID(id int) error {
	return repository.DeleteTaskByID(id)
}

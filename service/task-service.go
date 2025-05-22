package service

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/repository"
)

func CreateTask(task *model.Task) error {
	return repository.CreateTask(task)
}

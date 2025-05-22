package service

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/repository"
)

func CreateProject(project *model.Project) error {
	return repository.CreateProject(project)
}

package service

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/repository"
)

func CreateProject(project *model.Project) error {
	return repository.CreateProject(project)
}

func GetProjectByID(id int) ([]model.Project, error) {
	return repository.GetProjectsByUserID(id)
}

func GetProjectDetailByID(id int) (*model.Project, error) {
	return repository.GetProjectDetailByID(id)
}

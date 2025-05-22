package repository

import (
	"intervew-intern-dot/config"
	"intervew-intern-dot/model"
)

func CreateProject(project *model.Project) error {
	return config.DB.Create(project).Error
}

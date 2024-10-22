package database

import (
	"go-tasker/schemas"
	"go-tasker/types"
)

func (s *service) GetProjects() ([]schemas.Project, error) {
	var projects []schemas.Project
	if err := s.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *service) CreateProject(payload types.CreateProjectPayload) (*schemas.Project, error) {
	project := schemas.Project{
		Title:  payload.Title,
		Status: payload.Status,
	}

	if err := s.db.Create(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *service) UpdateProject(projectID string, payload types.UpdateProjectPayload) (*schemas.Project, error) {
	var project schemas.Project
	if err := s.db.First(&project, projectID).Error; err != nil {
		return nil, err
	}

	project.Title = payload.Title
	project.Status = payload.Status

	if err := s.db.Save(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *service) DeleteProject(projectID string) error {
	var project schemas.Project
	if err := s.db.First(&project, projectID).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&project).Error; err != nil {
		return err
	}

	return nil
}

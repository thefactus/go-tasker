package database

import (
	"go-tasker/schemas"
	"go-tasker/types"
	"strconv"
)

func (s *service) GetLists(projectID string) ([]schemas.List, error) {
	var lists []schemas.List
	if err := s.db.Where("project_id = ?", projectID).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (s *service) GetList(projectID string, listID string) (*schemas.List, error) {
	var list schemas.List
	if err := s.db.Where("id = ? AND project_id = ?", listID, projectID).First(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *service) CreateList(projectID string, payload types.CreateListPayload) (*schemas.List, error) {
	projectIDUint, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		return nil, err
	}

	list := schemas.List{
		Title:     payload.Title,
		ProjectID: uint(projectIDUint),
	}

	if err := s.db.Create(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (s *service) UpdateList(projectID string, listID string, payload types.UpdateListPayload) (*schemas.List, error) {
	var list schemas.List
	if err := s.db.Where("id = ? AND project_id = ?", listID, projectID).First(&list).Error; err != nil {
		return nil, err
	}

	list.Title = payload.Title

	if err := s.db.Save(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (s *service) DeleteList(projectID string, listID string) error {
	var list schemas.List
	if err := s.db.Where("id = ? AND project_id = ?", listID, projectID).First(&list).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

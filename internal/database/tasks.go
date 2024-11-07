package database

import (
	"go-tasker/schemas"
	"go-tasker/types"
	"strconv"
)

func (s *service) GetTasks(projectID string, listID string) ([]schemas.Task, error) {
	var list schemas.List
	if err := s.db.Where("id = ? AND project_id = ?", listID, projectID).First(&list).Error; err != nil {
		return nil, err
	}

	var tasks []schemas.Task
	if err := s.db.Where("list_id = ?", listID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *service) CreateTask(projectID string, listID string, payload types.CreateTaskPayload) (*schemas.Task, error) {
	var list schemas.List
	if err := s.db.Where("id = ? AND project_id = ?", listID, projectID).First(&list).Error; err != nil {
		return nil, err
	}

	listIDUint, err := strconv.ParseUint(listID, 10, 64)
	if err != nil {
		return nil, err
	}

	task := schemas.Task{
		Title:  payload.Title,
		ListID: uint(listIDUint),
	}

	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) UpdateTask(projectID string, listID string, taskID string, payload types.UpdateTaskPayload) (*schemas.Task, error) {
	var task schemas.Task
	if err := s.db.Joins("JOIN lists ON lists.id = tasks.list_id").
		Where("tasks.id = ? AND tasks.list_id = ? AND lists.project_id = ?", taskID, listID, projectID).
		First(&task).Error; err != nil {
		return nil, err
	}

	task.Title = payload.Title
	task.Done = payload.Done

	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) UpdateTaskDone(projectID string, listID string, taskID string, payload types.UpdateTaskDonePayload) (*schemas.Task, error) {
	var task schemas.Task
	if err := s.db.Joins("JOIN lists ON lists.id = tasks.list_id").
		Where("tasks.id = ? AND tasks.list_id = ? AND lists.project_id = ?", taskID, listID, projectID).
		First(&task).Error; err != nil {
		return nil, err
	}

	task.Done = payload.Done

	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) DeleteTask(projectID string, listID string, taskID string) error {
	var task schemas.Task
	if err := s.db.Joins("JOIN lists ON lists.id = tasks.list_id").
		Where("tasks.id = ? AND tasks.list_id = ? AND lists.project_id = ?", taskID, listID, projectID).
		First(&task).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}

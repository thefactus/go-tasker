package database

import (
	"strconv"
	"todolist/schemas"
	"todolist/types"
)

func (s *service) GetTasks(listID string) ([]schemas.Task, error) {
	var tasks []schemas.Task
	if err := s.db.Where("list_id = ?", listID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *service) CreateTask(listID string, payload types.CreateTaskPayload) (*schemas.Task, error) {
	listIDUint, err := strconv.ParseUint(listID, 10, 64)
	if err != nil {
		return nil, err
	}

	task := schemas.Task{
		Title:  payload.Title,
		ListID: uint(listIDUint), // Convert listID to uint type
	}

	// Create the task in the database
	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) UpdateTask(taskID string, payload types.UpdateTaskPayload) (*schemas.Task, error) {
	var task schemas.Task
	if err := s.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	task.Title = payload.Title
	task.Done = payload.Done

	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) UpdateTaskDone(taskID string, payload types.UpdateTaskDonePayload) (*schemas.Task, error) {
	var task schemas.Task
	if err := s.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}

	task.Done = payload.Done

	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *service) DeleteTask(taskID string) error {
	var task schemas.Task
	if err := s.db.First(&task, taskID).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}

package server

import (
	"errors"
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"

	"gorm.io/gorm"
)

// GetTasksHandler godoc
// @Summary Get all tasks for a list within a project
// @Description Get all tasks for a list within a project
// @Tags tasks
// @Produce json
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks [get]
func (s *Server) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")

	tasks, err := s.db.GetTasks(projectID, listID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "List not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Tasks retrieved successfully", tasks)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostTasksHandler godoc
// @Summary Create a new task for a list within a project
// @Description Create a new task for a list within a project
// @Tags tasks
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Param task body types.CreateTaskPayload true "Create Task Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks [post]
func (s *Server) PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")

	var createTaskPayload types.CreateTaskPayload
	if err := utils.ParseAndValidateJSON(w, r, &createTaskPayload); err != nil {
		return
	}

	task, err := s.db.CreateTask(projectID, listID, createTaskPayload)
	if err != nil {
		utils.WriteInternalServerError(w, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task created successfully", task)

	utils.WriteJSON(w, http.StatusCreated, response)
}

// PutTaskHandler godoc
// @Summary Update a task within a list and project
// @Description Update a task within a list and project
// @Tags tasks
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Param taskID path string true "Task ID"
// @Param task body types.UpdateTaskPayload true "Update Task Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID} [put]
func (s *Server) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")
	taskID := r.PathValue("taskID")

	var updateTaskPayload types.UpdateTaskPayload
	if err := utils.ParseAndValidateJSON(w, r, &updateTaskPayload); err != nil {
		return
	}

	task, err := s.db.UpdateTask(projectID, listID, taskID, updateTaskPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task updated successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteTaskHandler godoc
// @Summary Delete a task within a list and project
// @Description Delete a task within a list and project
// @Tags tasks
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string.interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID} [delete]
func (s *Server) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")
	taskID := r.PathValue("taskID")

	err := s.db.DeleteTask(projectID, listID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PatchTaskDoneHandler godoc
// @Summary Mark a task as done within a list and project
// @Description Mark a task as done within a list and project
// @Tags tasks
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string.interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}/done [patch]
func (s *Server) PatchTaskDoneHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")
	taskID := r.PathValue("taskID")

	var updateTaskDonePayload types.UpdateTaskDonePayload
	updateTaskDonePayload.Done = true

	task, err := s.db.UpdateTaskDone(projectID, listID, taskID, updateTaskDonePayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task marked as done successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PatchTaskUndoneHandler godoc
// @Summary Mark a task as undone within a list and project
// @Description Mark a task as undone within a list and project
// @Tags tasks
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param listID path string true "List ID"
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string.interface{}
// @Failure 400 {object} map[string.interface{}
// @Failure 500 {object} map[string.interface{}
// @Router /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}/undone [patch]
func (s *Server) PatchTaskUndoneHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("listID")
	taskID := r.PathValue("taskID")

	var updateTaskDonePayload types.UpdateTaskDonePayload
	updateTaskDonePayload.Done = false

	task, err := s.db.UpdateTaskDone(projectID, listID, taskID, updateTaskDonePayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task marked as undone successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

package server

import (
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"
)

// GetTasksHandler godoc
// @Summary Get all tasks for a list
// @Description Get all tasks for a list
// @Tags tasks
// @Produce json
// @Param listID path string true "List ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks [get]
func (s *Server) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("listID")
	tasks, err := s.db.GetTasks(taskID)
	if err != nil {
		http.Error(w, "Error getting tasks", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Tasks retrieved successfully", tasks)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostTasksHandler godoc
// @Summary Create a new task for a list
// @Description Create a new task for a list
// @Tags tasks
// @Accept json
// @Produce json
// @Param listID path string true "List ID"
// @Param task body types.CreateTaskPayload true "Create Task Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks [post]
func (s *Server) PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("listID")

	var createTaskPayload types.CreateTaskPayload
	if err := utils.ParseAndValidateJSON(w, r, &createTaskPayload); err != nil {
		return
	}

	task, err := s.db.CreateTask(taskID, createTaskPayload)
	if err != nil {
		utils.WriteInternalServerError(w, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task created successfully", task)

	utils.WriteJSON(w, http.StatusCreated, response)
}

// PutTaskHandler godoc
// @Summary Update a task
// @Description Update a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Param task body types.UpdateTaskPayload true "Update Task Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks/{taskID} [put]
func (s *Server) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("taskID")

	var updateTaskPayload types.UpdateTaskPayload
	if err := utils.ParseAndValidateJSON(w, r, &updateTaskPayload); err != nil {
		return
	}

	task, err := s.db.UpdateTask(taskID, updateTaskPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task updated successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteTaskHandler godoc
// @Summary Delete a task
// @Description Delete a task
// @Tags tasks
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks/{taskID} [delete]
func (s *Server) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("taskID")

	err := s.db.DeleteTask(taskID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PatchTaskDoneHandler godoc
// @Summary Mark a task as done
// @Description Mark a task as done
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks/{taskID}/done [patch]
func (s *Server) PatchTaskDoneHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("taskID")

	var updateTaskDonePayload types.UpdateTaskDonePayload
	updateTaskDonePayload.Done = true

	task, err := s.db.UpdateTaskDone(taskID, updateTaskDonePayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task marked as done successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PatchTaskUndoneHandler godoc
// @Summary Mark a task as undone
// @Description Mark a task as undone
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{listID}/tasks/{taskID}/undone [patch]
func (s *Server) PatchTaskUndoneHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("taskID")

	var updateTaskDonePayload types.UpdateTaskDonePayload
	updateTaskDonePayload.Done = false

	task, err := s.db.UpdateTaskDone(taskID, updateTaskDonePayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Task marked as undone successfully", task)

	utils.WriteJSON(w, http.StatusOK, response)
}

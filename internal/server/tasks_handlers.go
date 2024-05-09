package server

import (
	"net/http"
	"todolist/utils"
)

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

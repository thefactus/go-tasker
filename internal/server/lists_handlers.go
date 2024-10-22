package server

import (
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"
)

// GetListsHandler godoc
// @Summary Get all lists
// @Description Get all lists
// @Tags lists
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists [get]
func (s *Server) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := s.db.GetLists()
	if err != nil {
		http.Error(w, "Error getting lists", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Lists retrieved successfully", lists)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostListsHandler godoc
// @Summary Create a new list
// @Description Create a new list
// @Tags lists
// @Accept json
// @Produce json
// @Param list body types.CreateListPayload true "Create List Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists [post]
func (s *Server) PostListsHandler(w http.ResponseWriter, r *http.Request) {
	var createListPayload types.CreateListPayload
	if err := utils.ParseAndValidateJSON(w, r, &createListPayload); err != nil {
		return
	}

	list, err := s.db.CreateList(createListPayload)
	if err != nil {
		utils.WriteInternalServerError(w, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List created successfully", list)

	utils.WriteJSON(w, http.StatusCreated, response)
}

// PutListHandler godoc
// @Summary Update a list
// @Description Update a list
// @Tags lists
// @Accept json
// @Produce json
// @Param id path string true "List ID"
// @Param list body types.UpdateListPayload true "Update List Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{id} [put]
func (s *Server) PutListHandler(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("id")

	// Parse the request body
	var updateListPayload types.UpdateListPayload
	if err := utils.ParseAndValidateJSON(w, r, &updateListPayload); err != nil {
		return
	}

	// Update the list
	list, err := s.db.UpdateList(listID, updateListPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List updated successfully", list)

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteListHandler godoc
// @Summary Delete a list
// @Description Delete a list
// @Tags lists
// @Param id path string true "List ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/lists/{id} [delete]
func (s *Server) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("id")

	err := s.db.DeleteList(listID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetListsByProjectHandler godoc
// @Summary Get all lists by project
// @Description Get all lists by project
// @Tags lists
// @Produce json
// @Param projectID path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists [get]
func (s *Server) GetListsByProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	lists, err := s.db.GetListsByProject(projectID)
	if err != nil {
		http.Error(w, "Error getting lists by project", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Lists retrieved successfully", lists)

	utils.WriteJSON(w, http.StatusOK, response)
}

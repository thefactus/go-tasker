package server

import (
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"
)

// GetListsHandler godoc
// @Summary Get all lists for a project
// @Description Get all lists for a project
// @Tags lists
// @Produce json
// @Param projectID path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists [get]
func (s *Server) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	lists, err := s.db.GetLists(projectID)
	if err != nil {
		http.Error(w, "Error getting lists", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Lists retrieved successfully", lists)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostListsHandler godoc
// @Summary Create a new list within a project
// @Description Create a new list within a project
// @Tags lists
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param list body types.CreateListPayload true "Create List Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists [post]
func (s *Server) PostListsHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	var createListPayload types.CreateListPayload
	if err := utils.ParseAndValidateJSON(w, r, &createListPayload); err != nil {
		return
	}

	list, err := s.db.CreateList(projectID, createListPayload)
	if err != nil {
		utils.WriteInternalServerError(w, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List created successfully", list)

	utils.WriteJSON(w, http.StatusCreated, response)
}

// PutListHandler godoc
// @Summary Update a list within a project
// @Description Update a list within a project
// @Tags lists
// @Accept json
// @Produce json
// @Param projectID path string true "Project ID"
// @Param id path string true "List ID"
// @Param list body types.UpdateListPayload true "Update List Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{id} [put]
func (s *Server) PutListHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("id")

	var updateListPayload types.UpdateListPayload
	if err := utils.ParseAndValidateJSON(w, r, &updateListPayload); err != nil {
		return
	}

	list, err := s.db.UpdateList(projectID, listID, updateListPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List updated successfully", list)

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteListHandler godoc
// @Summary Delete a list within a project
// @Description Delete a list within a project
// @Tags lists
// @Param projectID path string true "Project ID"
// @Param id path string true "List ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{id} [delete]
func (s *Server) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("id")

	err := s.db.DeleteList(projectID, listID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetListHandler godoc
// @Summary Get a specific list within a project
// @Description Get a specific list within a project
// @Tags lists
// @Produce json
// @Param projectID path string true "Project ID"
// @Param id path string true "List ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{projectID}/lists/{id} [get]
func (s *Server) GetListHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")
	listID := r.PathValue("id")

	list, err := s.db.GetList(projectID, listID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List retrieved successfully", list)

	utils.WriteJSON(w, http.StatusOK, response)
}

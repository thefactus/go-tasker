package server

import (
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"
)

func (s *Server) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := s.db.GetLists()
	if err != nil {
		http.Error(w, "Error getting lists", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Lists retrieved successfully", lists)

	utils.WriteJSON(w, http.StatusOK, response)
}

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
		// maybe change this error
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List updated successfully", list)

	utils.WriteJSON(w, http.StatusOK, response)
}

func (s *Server) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("id")

	err := s.db.DeleteList(listID)
	if err != nil {
		// maybe change this error
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("List deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

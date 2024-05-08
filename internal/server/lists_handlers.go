package server

import (
	"net/http"
	"todolist/types"
	"todolist/utils"
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

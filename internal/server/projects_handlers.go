package server

import (
	"go-tasker/types"
	"go-tasker/utils"
	"net/http"
)

// GetProjectsHandler godoc
// @Summary Get all projects
// @Description Get all projects
// @Tags projects
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects [get]
func (s *Server) GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := s.db.GetProjects()
	if err != nil {
		http.Error(w, "Error getting projects", http.StatusInternalServerError)
		return
	}

	response := utils.PrepareJSONWithMessage("Projects retrieved successfully", projects)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostProjectsHandler godoc
// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Param project body types.CreateProjectPayload true "Create Project Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects [post]
func (s *Server) PostProjectsHandler(w http.ResponseWriter, r *http.Request) {
	var createProjectPayload types.CreateProjectPayload
	if err := utils.ParseAndValidateJSON(w, r, &createProjectPayload); err != nil {
		return
	}

	project, err := s.db.CreateProject(createProjectPayload)
	if err != nil {
		utils.WriteInternalServerError(w, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Project created successfully", project)

	utils.WriteJSON(w, http.StatusCreated, response)
}

// PutProjectHandler godoc
// @Summary Update a project
// @Description Update a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body types.UpdateProjectPayload true "Update Project Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{id} [put]
func (s *Server) PutProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("id")

	var updateProjectPayload types.UpdateProjectPayload
	if err := utils.ParseAndValidateJSON(w, r, &updateProjectPayload); err != nil {
		return
	}

	project, err := s.db.UpdateProject(projectID, updateProjectPayload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Project updated successfully", project)

	utils.WriteJSON(w, http.StatusOK, response)
}

// DeleteProjectHandler godoc
// @Summary Delete a project
// @Description Delete a project
// @Tags projects
// @Param id path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/projects/{id} [delete]
func (s *Server) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("id")

	err := s.db.DeleteProject(projectID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := utils.PrepareJSONWithMessage("Project deleted successfully", nil)

	utils.WriteJSON(w, http.StatusOK, response)
}

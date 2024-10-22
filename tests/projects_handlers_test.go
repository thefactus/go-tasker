package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestProjects(t *testing.T) {
	t.Run("expects to get empty data", func(t *testing.T) {
		clearTableProjects()

		req, _ := http.NewRequest("GET", "/api/v1/projects", nil)
		response := executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		expected := Response{
			Data:    []interface{}{},
			Message: "Projects retrieved successfully",
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected '%v'. Got '%v'", expected, actual)
		}
	})

	t.Run("expects to get a project", func(t *testing.T) {
		clearTableProjects()

		payload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		req, _ = http.NewRequest("GET", "/api/v1/projects", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		for _, item := range actual.Data {
			project, ok := item.(map[string]interface{})
			if !ok {
				t.Errorf("Expected project item to be a map. Got '%v'", item)
			}

			assert.Equal(t, float64(1), project["id"], "Expected id to be 1")
			assert.Equal(t, "Project 1", project["title"], "Expected title to be 'Project 1'")
			assert.Equal(t, "not started", project["status"], "Expected status to be 'not started'")
			assert.NotNil(t, project["created_at"], "Expected 'created_at' field to be present")
		}
	})

	t.Run("expects to create a project", func(t *testing.T) {
		clearTableProjects()

		payload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "Project created successfully", result["message"],
			"Expected message to be 'Project created successfully'")
		assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")
		assert.Equal(t, "Project 1", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Project 1'")
		assert.Equal(t, "not started", result["data"].(map[string]interface{})["status"],
			"Expected status to be 'not started'")
		assert.NotNil(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
	})

	t.Run("while creating/when title is missing/expects to return validation error", func(t *testing.T) {
		clearTableProjects()

		payload := []byte(`{"status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, result["error"], "Missing required fields: title",
			"Expected error to be 'Missing required fields: title'")
	})

	t.Run("expects to update a project", func(t *testing.T) {
		clearTableProjects()

		payload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		payload = []byte(`{"title": "Project 1 Updated", "status": "in progress"}`)
		req, _ = http.NewRequest("PUT", "/api/v1/projects/1", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "Project updated successfully", result["message"],
			"Expected message to be 'Project updated successfully'")
		assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")
		assert.Equal(t, "Project 1 Updated", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Project 1 Updated'")
		assert.Equal(t, "in progress", result["data"].(map[string]interface{})["status"],
			"Expected status to be 'in progress'")
		assert.NotNil(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
	})

	t.Run("expects to delete a project", func(t *testing.T) {
		clearTableProjects()

		payload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		req, _ = http.NewRequest("DELETE", "/api/v1/projects/1", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		// Check for message
		assert.Equal(t, result["message"], "Project deleted successfully",
			"Expected message to be 'Project deleted successfully'")
	})
}

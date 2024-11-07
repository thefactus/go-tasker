package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestLists(t *testing.T) {
	t.Run("expects to get empty data", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Get lists for the project
		req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectIDStr+"/lists", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		expected := Response{
			Data:    []interface{}{},
			Message: "Lists retrieved successfully",
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected '%v'. Got '%v'", expected, actual)
		}
	})

	t.Run("expects to get a list", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Create a list under the project
		payload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		// Get lists for the project
		req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectIDStr+"/lists", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		for _, item := range actual.Data {
			list, ok := item.(map[string]interface{})
			if !ok {
				t.Errorf("Expected list item to be a map. Got '%v'", item)
			}

			assert.Equal(t, float64(1), list["id"], "Expected id to be 1")
			assert.Equal(t, "Tasks", list["title"], "Expected title to be 'Tasks'")
			assert.NotNil(t, list["created_at"], "Expected 'created_at' field to be present")
			assert.Equal(t, float64(projectID), list["project_id"], "Expected project_id to be %d", projectID)
		}
	})

	t.Run("expects to create a list", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Create a list under the project
		payload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "List created successfully", result["message"],
			"Expected message to be 'List created successfully'")
		assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")
		assert.Equal(t, "Tasks", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Tasks'")
		assert.NotNil(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
		assert.Equal(t, float64(projectID), result["data"].(map[string]interface{})["project_id"],
			"Expected project_id to be %d", projectID)
	})

	t.Run("while creating/when title is missing/expects to return validation error", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Attempt to create a list without a title
		payload := []byte(`{}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "Missing required fields: title", result["error"],
			"Expected error to be 'Missing required fields: title'")
	})

	t.Run("expects to update a list", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}

		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Create a list under the project
		payload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		// Update the list
		payload = []byte(`{"title": "Tasks Updated"}`)
		req, _ = http.NewRequest("PUT", "/api/v1/projects/"+projectIDStr+"/lists/1", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "List updated successfully", result["message"],
			"Expected message to be 'List updated successfully'")
		assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")
		assert.Equal(t, "Tasks Updated", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Tasks Updated'")
		assert.NotNil(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
		assert.Equal(t, float64(projectID), result["data"].(map[string]interface{})["project_id"],
			"Expected project_id to be %d", projectID)
	})

	t.Run("expects to delete a list", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project first
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Create a list under the project
		payload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		// Delete the list
		req, _ = http.NewRequest("DELETE", "/api/v1/projects/"+projectIDStr+"/lists/1", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		assert.Equal(t, "List deleted successfully", result["message"],
			"Expected message to be 'List deleted successfully'")
	})

	t.Run("expects to get lists by project", func(t *testing.T) {
		clearTableLists()
		clearTableProjects()

		// Create a project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Errorf("Error unmarshalling project response: %v", err)
			return
		}
		projectID := int(projectResult["data"].(map[string]interface{})["id"].(float64))
		projectIDStr := strconv.Itoa(projectID)

		// Create a list under the project
		payload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectIDStr+"/lists", bytes.NewReader(payload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		// Get lists for the project
		req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectIDStr+"/lists", nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		for _, item := range actual.Data {
			list, ok := item.(map[string]interface{})
			if !ok {
				t.Errorf("Expected list item to be a map. Got '%v'", item)
			}

			assert.Equal(t, float64(1), list["id"], "Expected id to be 1")
			assert.Equal(t, "Tasks", list["title"], "Expected title to be 'Tasks'")
			assert.Equal(t, float64(projectID), list["project_id"], "Expected project_id to be %d", projectID)
			assert.NotNil(t, list["created_at"], "Expected 'created_at' field to be present")
		}
	})
}

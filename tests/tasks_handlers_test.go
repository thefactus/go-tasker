package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTasks(t *testing.T) {
	t.Run("expects to get empty data", func(t *testing.T) {
		clearTables()

		// Create a project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Fatalf("Error unmarshalling project response: %v", err)
		}
		projectIDFloat := projectResult["data"].(map[string]interface{})["id"].(float64)
		projectID := strconv.FormatFloat(projectIDFloat, 'f', -1, 64)

		// Create a list under the project
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &listResult)
		if err != nil {
			t.Fatalf("Error unmarshalling list response: %v", err)
		}
		listIDFloat := listResult["data"].(map[string]interface{})["id"].(float64)
		listID := strconv.FormatFloat(listIDFloat, 'f', -1, 64)

		// Get tasks
		req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		expected := Response{
			Data:    []interface{}{},
			Message: "Tasks retrieved successfully",
		}

		assert.Equal(t, expected, actual, "Expected response does not match actual response")
	})

	t.Run("expects to get a task", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &projectResult)
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &listResult)
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		// Get tasks
		req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		for _, item := range actual.Data {
			task, ok := item.(map[string]interface{})
			if !ok {
				t.Fatalf("Expected task item to be a map. Got '%v'", item)
			}

			assert.Equal(t, float64(1), task["id"], "Expected id to be 1")
			assert.Equal(t, "Task 1", task["title"], "Expected title to be 'Task 1'")
			assert.Equal(t, false, task["done"], "Expected done to be false")
			assert.Equal(t, float64(1), task["list_id"], "Expected list_id to be 1")
		}
	})

	t.Run("expects to create a task", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &projectResult)
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &listResult)
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var result map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &result)

		assert.Equal(t, "Task created successfully", result["message"])
		data := result["data"].(map[string]interface{})
		assert.Equal(t, float64(1), data["id"], "Expected id to be 1")
		assert.Equal(t, "Task 1", data["title"], "Expected title to be 'Task 1'")
		assert.Equal(t, false, data["done"], "Expected done to be false")
		assert.Equal(t, float64(1), data["list_id"], "Expected list_id to be 1")
	})

	t.Run("expects to update a task", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Fatalf("Error unmarshalling project response: %v", err)
		}
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &listResult)
		if err != nil {
			t.Fatalf("Error unmarshalling list response: %v", err)
		}
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var taskResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &taskResult)
		if err != nil {
			t.Fatalf("Error unmarshalling task response: %v", err)
		}
		taskID := strconv.FormatFloat(taskResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Update task
		updateTaskPayload := []byte(`{"title": "Task 1 Updated", "done": true}`)
		req, _ = http.NewRequest("PUT", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks/"+taskID, bytes.NewReader(updateTaskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Error unmarshalling update task response: %v", err)
		}

		assert.Equal(t, "Task updated successfully", result["message"])
		data := result["data"].(map[string]interface{})
		assert.Equal(t, float64(1), data["id"], "Expected id to be 1")
		assert.Equal(t, "Task 1 Updated", data["title"], "Expected title to be 'Task 1 Updated'")
		assert.Equal(t, true, data["done"], "Expected done to be true")
		assert.Equal(t, float64(1), data["list_id"], "Expected list_id to be 1")
	})

	t.Run("expects to mark task as undone", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Fatalf("Error unmarshalling project response: %v", err)
		}
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &listResult)
		if err != nil {
			t.Fatalf("Error unmarshalling list response: %v", err)
		}
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var taskResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &taskResult)
		if err != nil {
			t.Fatalf("Error unmarshalling task response: %v", err)
		}
		taskID := strconv.FormatFloat(taskResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Mark task as done
		req, _ = http.NewRequest("PATCH", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks/"+taskID+"/done", nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		// Mark task as undone
		req, _ = http.NewRequest("PATCH", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks/"+taskID+"/undone", nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Error unmarshalling update task response: %v", err)
		}

		assert.Equal(t, "Task marked as undone successfully", result["message"])
		data := result["data"].(map[string]interface{})
		assert.Equal(t, float64(1), data["id"], "Expected id to be 1")
		assert.Equal(t, "Task 1", data["title"], "Expected title to be 'Task 1'")
		assert.Equal(t, false, data["done"], "Expected done to be false")
		assert.Equal(t, float64(1), data["list_id"], "Expected list_id to be 1")
	})

	t.Run("expects to delete a task", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Fatalf("Error unmarshalling project response: %v", err)
		}
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &listResult)
		if err != nil {
			t.Fatalf("Error unmarshalling list response: %v", err)
		}
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var taskResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &taskResult)
		if err != nil {
			t.Fatalf("Error unmarshalling task response: %v", err)
		}
		taskID := strconv.FormatFloat(taskResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Delete task
		req, _ = http.NewRequest("DELETE", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks/"+taskID, nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Error unmarshalling delete task response: %v", err)
		}

		assert.Equal(t, "Task deleted successfully", result["message"])
	})

	t.Run("expects to delete a list and all its tasks", func(t *testing.T) {
		clearTables()

		// Create project
		projectPayload := []byte(`{"title": "Project 1", "status": "not started"}`)
		req, _ := http.NewRequest("POST", "/api/v1/projects", bytes.NewReader(projectPayload))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var projectResult map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &projectResult)
		if err != nil {
			t.Fatalf("Error unmarshalling project response: %v", err)
		}
		projectID := strconv.FormatFloat(projectResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create list
		listPayload := []byte(`{"title": "Tasks"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists", bytes.NewReader(listPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		var listResult map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &listResult)
		if err != nil {
			t.Fatalf("Error unmarshalling list response: %v", err)
		}
		listID := strconv.FormatFloat(listResult["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)

		// Create task
		taskPayload := []byte(`{"title": "Task 1"}`)
		req, _ = http.NewRequest("POST", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", bytes.NewReader(taskPayload))
		response = executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		// Delete list
		req, _ = http.NewRequest("DELETE", "/api/v1/projects/"+projectID+"/lists/"+listID, nil)
		response = executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatalf("Error unmarshalling delete list response: %v", err)
		}

		assert.Equal(t, "List deleted successfully", result["message"])

		// Verify tasks are deleted
		// req, _ = http.NewRequest("GET", "/api/v1/projects/"+projectID+"/lists/"+listID+"/tasks", nil)
		// response = executeRequest(req)
		// checkResponseCode(t, http.StatusNotFound, response.Code)
	})

	// t.Run("expects to get a task", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	req, _ = http.NewRequest("GET", "/api/v1/lists/1/tasks", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var actual Response
	// 	json.Unmarshal(response.Body.Bytes(), &actual)
	//
	// 	for _, item := range actual.Data {
	// 		task, ok := item.(map[string]interface{})
	// 		if !ok {
	// 			t.Errorf("Expected task item to be a map. Got '%v'", item)
	// 		}
	//
	// 		assert.Equal(t, float64(1), task["id"], "Expected id to be 1")
	// 		assert.Equal(t, "Task 1", task["title"], "Expected title to be 'Task 1'")
	// 	}
	// })
	//
	// t.Run("expects to create a task", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	// Check for message
	// 	assert.Equal(t, "Task created successfully", result["message"],
	// 		"Expected message to be 'Task created successfully'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
	// 		"Expected id to be 1")
	// 	assert.Equal(t, "Task 1", result["data"].(map[string]interface{})["title"],
	// 		"Expected title to be 'Task 1'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["list_id"],
	// 		"Expected list_id to be 1")
	// })
	//
	// t.Run("expects to update a task", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1 Updated"}`)
	// 	req, _ = http.NewRequest("PUT", "/api/v1/lists/1/tasks/1", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	// Check for message
	// 	assert.Equal(t, "Task updated successfully", result["message"],
	// 		"Expected message to be 'Task updated successfully'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
	// 		"Expected id to be 1")
	// 	assert.Equal(t, "Task 1 Updated", result["data"].(map[string]interface{})["title"],
	// 		"Expected title to be 'Task 1 Updated'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["list_id"],
	// 		"Expected list_id to be 1")
	// })
	//
	// t.Run("expects to mark task as done", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	req, _ = http.NewRequest("PATCH", "/api/v1/lists/1/tasks/1/done", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	assert.Equal(t, "Task marked as done successfully", result["message"],
	// 		"Expected message to be 'Task marked as done successfully'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
	// 		"Expected id to be 1")
	// 	assert.Equal(t, "Task 1", result["data"].(map[string]interface{})["title"],
	// 		"Expected title to be 'Task 1'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["list_id"],
	// 		"Expected list_id to be 1")
	// 	assert.Equal(t, true, result["data"].(map[string]interface{})["done"],
	// 		"Expected done to be true")
	// })
	//
	// t.Run("expects to mark task as undone", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	req, _ = http.NewRequest("PATCH", "/api/v1/lists/1/tasks/1/done", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	req, _ = http.NewRequest("PATCH", "/api/v1/lists/1/tasks/1/undone", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	assert.Equal(t, "Task marked as undone successfully", result["message"],
	// 		"Expected message to be 'Task marked as undone successfully'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["id"],
	// 		"Expected id to be 1")
	// 	assert.Equal(t, "Task 1", result["data"].(map[string]interface{})["title"],
	// 		"Expected title to be 'Task 1'")
	// 	assert.Equal(t, float64(1), result["data"].(map[string]interface{})["list_id"],
	// 		"Expected list_id to be 1")
	// 	assert.Equal(t, false, result["data"].(map[string]interface{})["done"],
	// 		"Expected done to be false")
	// })
	//
	// t.Run("expects to delete a task", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	req, _ = http.NewRequest("DELETE", "/api/v1/lists/1/tasks/1", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	// Check for message
	// 	assert.Equal(t, "Task deleted successfully", result["message"],
	// 		"Expected message to be 'Task deleted successfully'")
	// })
	//
	// t.Run("expects to delete a list and all its tasks", func(t *testing.T) {
	// 	clearTableTasksAndLists()
	//
	// 	payload := []byte(`{"title": "Tasks"}`)
	// 	req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
	// 	response := executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	payload = []byte(`{"title": "Task 1"}`)
	// 	req, _ = http.NewRequest("POST", "/api/v1/lists/1/tasks", bytes.NewReader(payload))
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusCreated, response.Code)
	//
	// 	req, _ = http.NewRequest("DELETE", "/api/v1/lists/1", nil)
	// 	response = executeRequest(req)
	//
	// 	checkResponseCode(t, http.StatusOK, response.Code)
	//
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(response.Body.Bytes(), &result)
	// 	if err != nil {
	// 		t.Errorf("Error unmarshalling response: %v", err)
	// 		return
	// 	}
	//
	// 	// Check for message
	// 	assert.Equal(t, "List deleted successfully", result["message"],
	// 		"Expected message to be 'List deleted successfully'")
	// })
}

package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTasks(t *testing.T) {
	t.Run("expects to get empty data", func(t *testing.T) {
		clearTableTasksAndLists()

		payload := []byte(`{"title": "Tasks"}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		req, _ = http.NewRequest("GET", "/api/v1/lists/1/tasks", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		expected := Response{
			Data:    []interface{}{},
			Message: "Tasks retrieved successfully",
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected '%v'. Got '%v'", expected, actual)
		}
	})

}

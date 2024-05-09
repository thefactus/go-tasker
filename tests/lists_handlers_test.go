package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"todolist/internal/server"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var (
	s  *http.Server
	db *gorm.DB
)

type Response struct {
	Data    []interface{} `json:"data"`
	Message string        `json:"message"`
}

func TestMain(m *testing.M) {
	s = server.NewServer()
	db = getDB()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func getDB() *gorm.DB {
	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func clearTable() {
	db.Exec("DELETE FROM lists")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='lists'") // sqlite3
	// db.Exec("ALTER SEQUENCE lists_id_seq RESTART WITH 1")  // postgres
	// db.Exec("ALTER TABLE lists AUTO_INCREMENT = 1")        // mysql
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Handler.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestLists(t *testing.T) {
	t.Run("expects to get empty data", func(t *testing.T) {
		clearTable()

		req, _ := http.NewRequest("GET", "/api/v1/lists", nil)
		response := executeRequest(req)

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
		clearTable()

		payload := []byte(`{"title": "Tasks"}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		req, _ = http.NewRequest("GET", "/api/v1/lists", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var actual Response
		json.Unmarshal(response.Body.Bytes(), &actual)

		for _, item := range actual.Data {
			list, ok := item.(map[string]interface{})
			if !ok {
				t.Errorf("Expected list item to be a map. Got '%v'", item)
			}

			assertEqual(t, float64(1), list["id"], "Expected id to be 1")

			assertEqual(t, "Tasks", list["title"], "Expected title to be 'Tasks'")

			assertIsPresent(t, list["created_at"], "Expected 'created_at' field to be present")
		}
	})

	t.Run("expects to create a list", func(t *testing.T) {
		clearTable()

		payload := []byte(`{"title": "Tasks"}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		// Check for message
		assertEqual(t, "List created successfully", result["message"],
			"Expected message to be 'List created successfully'")

		// Check for id
		assertEqual(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")

		// Check for title
		assertEqual(t, "Tasks", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Tasks'")

		// Check for CreatedAt presence
		assertIsPresent(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
	})

	t.Run("while creating/when title is missing/expects to return validation error", func(t *testing.T) {
		clearTable()

		payload := []byte(`{}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusBadRequest, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		// Check for error message
		assertEqual(t, result["error"], "Missing required fields: title",
			"Expected error to be 'Missing required fields: title'")
	})

	t.Run("expects to update a list", func(t *testing.T) {
		clearTable()

		payload := []byte(`{"title": "Tasks"}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		payload = []byte(`{"title": "Tasks Updated"}`)
		req, _ = http.NewRequest("PUT", "/api/v1/lists/1", bytes.NewReader(payload))
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		// Check for message
		assertEqual(t, "List updated successfully", result["message"],
			"Expected message to be 'List updated successfully'")

		// Check for id
		assertEqual(t, float64(1), result["data"].(map[string]interface{})["id"],
			"Expected id to be 1")

		// Check for title
		assertEqual(t, "Tasks Updated", result["data"].(map[string]interface{})["title"],
			"Expected title to be 'Tasks Updated'")

		// Check for CreatedAt presence
		assertIsPresent(t, result["data"].(map[string]interface{})["created_at"],
			"Expected 'created_at' field to be present")
	})

	t.Run("expects to delete a list", func(t *testing.T) {
		clearTable()

		payload := []byte(`{"title": "Tasks"}`)
		req, _ := http.NewRequest("POST", "/api/v1/lists", bytes.NewReader(payload))
		response := executeRequest(req)

		checkResponseCode(t, http.StatusCreated, response.Code)

		req, _ = http.NewRequest("DELETE", "/api/v1/lists/1", nil)
		response = executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err := json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}

		// Check for message
		assertEqual(t, result["message"], "List deleted successfully",
			"Expected message to be 'List deleted successfully'")
	})
}

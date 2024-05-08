package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestMain(m *testing.M) {
	s = server.NewServer()
	db = getDB()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func getDB() *gorm.DB {
	// Create DB and connect
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

func TestCreateList(t *testing.T) {
	clearTable()

	payload := []byte(`{"title": "Tasks"}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/v1/lists", bytes.NewReader(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var result map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	// Check for message
	if message, ok := result["message"].(string); !ok || message != "List created successfully" {
		t.Errorf("Expected message to be 'List created successfully'. Got %v", result["message"])
	}

	// Check for id
	if id, ok := result["data"].(map[string]interface{})["id"].(float64); !ok || id != 1 {
		t.Errorf("Expected id to be 1. Got %v", result["data"].(map[string]interface{})["id"])
	}

	// Check for title
	if title, ok := result["data"].(map[string]interface{})["title"].(string); !ok || title != "Tasks" {
		t.Errorf("Expected title to be 'Tasks'. Got %v", result["data"].(map[string]interface{})["title"])
	}

	// Check for CreatedAt presence
	if _, ok := result["data"].(map[string]interface{})["created_at"]; !ok {
		t.Errorf("Expected 'created_at' field to be present")
	}
}

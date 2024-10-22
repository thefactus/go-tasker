package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	s  *http.Server
	db *gorm.DB
)

type Response struct {
	Data    []interface{} `json:"data"`
	Message string        `json:"message"`
}

func clearTableLists() {
	db.Exec("DELETE FROM lists")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='lists'") // sqlite3
	// db.Exec("ALTER SEQUENCE lists_id_seq RESTART WITH 1")  // postgres
	// db.Exec("ALTER TABLE lists AUTO_INCREMENT = 1")        // mysql
}

func clearTableTasksAndLists() {
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM lists")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='tasks'") // sqlite3
	db.Exec("DELETE FROM sqlite_sequence WHERE name='lists'") // sqlite3
	// db.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")  // postgres
	// db.Exec("ALTER SEQUENCE lists_id_seq RESTART WITH 1")  // postgres
	// db.Exec("ALTER TABLE tasks AUTO_INCREMENT = 1")        // mysql
	// db.Exec("ALTER TABLE lists AUTO_INCREMENT = 1")        // mysql
}

func getDB() *gorm.DB {
	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
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

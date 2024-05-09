package database

import (
	"log"
	"os"
	"todolist/schemas"
	"todolist/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	GetLists() ([]schemas.List, error)
	CreateList(payload types.CreateListPayload) (*schemas.List, error)
	UpdateList(listID string, payload types.UpdateListPayload) (*schemas.List, error)
	DeleteList(listID string) error
	GetTasks(listID string) ([]schemas.Task, error)
}

type service struct {
	db *gorm.DB
}

var (
	dbUrl      = os.Getenv("DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.List{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&schemas.Task{})
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

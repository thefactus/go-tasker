package database

import (
	"go-tasker/schemas"
	"go-tasker/types"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	GetLists(projectID string) ([]schemas.List, error)
	GetList(projectID string, listID string) (*schemas.List, error)
	CreateList(projectID string, payload types.CreateListPayload) (*schemas.List, error)
	UpdateList(projectID string, listID string, payload types.UpdateListPayload) (*schemas.List, error)
	DeleteList(projectID string, listID string) error

	GetTasks(projectID string, listID string) ([]schemas.Task, error)
	CreateTask(projectID string, listID string, payload types.CreateTaskPayload) (*schemas.Task, error)
	UpdateTask(projectID string, listID string, taskID string, payload types.UpdateTaskPayload) (*schemas.Task, error)
	UpdateTaskDone(projectID string, listID string, taskID string, payload types.UpdateTaskDonePayload) (*schemas.Task, error)
	DeleteTask(projectID string, listID string, taskID string) error

	GetProjects() ([]schemas.Project, error)
	CreateProject(payload types.CreateProjectPayload) (*schemas.Project, error)
	UpdateProject(projectID string, payload types.UpdateProjectPayload) (*schemas.Project, error)
	DeleteProject(projectID string) error
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

	// Migrate the Schemas
	if err := db.AutoMigrate(&schemas.List{}); err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&schemas.Task{}); err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&schemas.Project{}); err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

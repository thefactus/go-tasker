package schemas

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title  string
	Done   bool
	ListID uint
}

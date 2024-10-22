package schemas

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title  string
	Status string
	Lists  []List
}

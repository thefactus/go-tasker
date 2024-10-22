package schemas

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title     string
	ProjectID uint
}

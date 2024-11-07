// Updated schemas/list.go
package schemas

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title     string
	ProjectID uint
	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tasks     []Task  `gorm:"constraint:OnDelete:CASCADE;"`
}

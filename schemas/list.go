package schemas

import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title string
}

type ListResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

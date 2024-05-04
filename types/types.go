package types

import "time"

type List struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateListPayload struct {
	Title string `json:"title" validate:"required"`
}

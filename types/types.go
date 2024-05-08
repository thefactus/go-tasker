package types

type CreateListPayload struct {
	Title string `json:"title" validate:"required"`
}

type UpdateListPayload struct {
	Title string `json:"title" validate:"required"`
}

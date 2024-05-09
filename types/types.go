package types

type CreateListPayload struct {
	Title string `json:"title" validate:"required"`
}

type UpdateListPayload struct {
	Title string `json:"title" validate:"required"`
}

type CreateTaskPayload struct {
	Title string `json:"title" validate:"required"`
}

type UpdateTaskPayload struct {
	Title string `json:"title" validate:"required"`
	Done  bool   `json:"done"`
}

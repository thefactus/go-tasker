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

type UpdateTaskDonePayload struct {
	Done bool `json:"done"`
}

type CreateProjectPayload struct {
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type UpdateProjectPayload struct {
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required"`
}

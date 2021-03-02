package dto

type CreateTodoRequest struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}

type CreateTodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

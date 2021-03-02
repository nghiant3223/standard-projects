package dto

type CreateTodoRequest struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}

type CreateTodoResponse struct {
	ID          int
	Title       string
	Description string
}

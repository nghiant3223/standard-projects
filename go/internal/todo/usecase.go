package todo

import (
	"context"

	"github.com/nghiant3223/standard-project/internal/todo/model"
)

//go:generate mockgen -source=usecase.go -destination=gomock/usecase.go
type UseCase interface {
	ListTodos(ctx context.Context) ([]model.Todo, error)
	GetTodoByID(ctx context.Context, id int) (model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

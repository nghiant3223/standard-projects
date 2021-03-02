package usecase

import (
	"context"

	"github.com/nghiant3223/standard-project/internal/todo"
	"github.com/nghiant3223/standard-project/internal/todo/model"
)

type UseCase struct {
	repository todo.Repository
}

func NewUseCase(repository todo.Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (uc *UseCase) ListTodos(context.Context) ([]model.Todo, error) {
	return uc.repository.List()
}

func (uc *UseCase) GetTodoByID(ctx context.Context, id int) (model.Todo, error) {
	return uc.repository.Get(id)
}

func (uc *UseCase) CreateTodo(ctx context.Context, todo *model.Todo) error {
	return uc.repository.Create(todo)
}

func (uc *UseCase) DeleteTodo(ctx context.Context, id int) error {
	return uc.repository.Delete(id)
}

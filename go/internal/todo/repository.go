package todo

import "github.com/nghiant3223/standard-project/internal/todo/model"

//go:generate mockgen -source=repository.go -destination=gomock/repository.go
type Repository interface {
	List() ([]model.Todo, error)
	Get(id int) (model.Todo, error)
	Create(todo *model.Todo) error
	Delete(id int) error
}

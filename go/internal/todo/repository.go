package todo

import "github.com/nghiant3223/standard-project/internal/todo/model"

type Repository interface {
	List() ([]model.Todo, error)
	Get(id int) (model.Todo, error)
	Create(todo *model.Todo) error
	Delete(id int) error
}

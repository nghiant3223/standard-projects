package repository

import (
	"errors"

	apperror "github.com/nghiant3223/standard-project/internal/todo/apperror"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"gorm.io/gorm"
)

type SQLRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) List() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *SQLRepository) Get(id int) (model.Todo, error) {
	var td model.Todo
	err := r.db.First(&td, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Todo{}, apperror.ErrNotFound
	}
	return td, err
}

func (r *SQLRepository) Create(todo *model.Todo) error {
	err := r.db.Create(todo).Error
	return err
}

func (r *SQLRepository) Delete(id int) error {
	err := r.db.Delete(&model.Todo{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.ErrNotFound
	}
	return err
}

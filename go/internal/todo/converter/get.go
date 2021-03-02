package converter

import (
	"github.com/nghiant3223/standard-project/internal/todo/dto"
	"github.com/nghiant3223/standard-project/internal/todo/model"
)

type GetTodoConverter struct{}

func NewGetTodoConverter() *GetTodoConverter {
	return &GetTodoConverter{}
}

func (t *GetTodoConverter) ToDto(td model.Todo) dto.CreateTodoResponse {
	return dto.CreateTodoResponse{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
	}
}

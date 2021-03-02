package converter

import (
	"github.com/nghiant3223/standard-project/internal/todo/dto"
	"github.com/nghiant3223/standard-project/internal/todo/model"
)

type CreateTodoConverter struct{}

func NewCreateTodoTransformer() *CreateTodoConverter {
	return &CreateTodoConverter{}
}

func (t *CreateTodoConverter) ToEntity(req dto.CreateTodoRequest) model.Todo {
	return model.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
}

func (t *CreateTodoConverter) ToDto(td model.Todo) dto.CreateTodoResponse {
	return dto.CreateTodoResponse{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
	}
}

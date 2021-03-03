package todo

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nghiant3223/standard-project/internal/todo"
	mocktodo "github.com/nghiant3223/standard-project/internal/todo/gomock"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/nghiant3223/standard-project/internal/todo/usecase"
	"github.com/stretchr/testify/suite"
)

type useCaseTestSuite struct {
	suite.Suite

	repo    *mocktodo.MockRepository
	useCase todo.UseCase
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, &useCaseTestSuite{})
}

func (s *useCaseTestSuite) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())
	s.repo = mocktodo.NewMockRepository(mockCtrl)
	s.useCase = usecase.New(s.repo)
}

func (s *useCaseTestSuite) Test_CreateTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Do it now!",
	}
	newTd := model.Todo{
		ID:          3930,
		Title:       td.Title,
		Description: td.Description,
	}

	s.repo.EXPECT().Create(&td).SetArg(0, newTd).Return(nil)

	err := s.useCase.CreateTodo(context.Background(), &td)
	s.NoError(err)
}

func (s *useCaseTestSuite) Test_GetTodo_Happy() {
	td := model.Todo{
		ID:          3930,
		Title:       "Clean the floor",
		Description: "Do it now!",
	}

	s.repo.EXPECT().Get(td.ID).Return(td, nil)

	result, err := s.useCase.GetTodoByID(context.Background(), td.ID)
	s.NoError(err)
	s.Equal(td, result)
}

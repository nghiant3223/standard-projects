package todo

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nghiant3223/standard-project/internal/todo"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/nghiant3223/standard-project/internal/todo/repository"
	"github.com/nghiant3223/standard-project/internal/todo/usecase"
	"github.com/nghiant3223/standard-project/pkg/configurator"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type useCaseTestSuite struct {
	suite.Suite

	dbConnection *gorm.DB
	dbMigration  *migrate.Migrate
	useCase      todo.UseCase
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, &useCaseTestSuite{})
}

func (s *useCaseTestSuite) SetupSuite() {
	configurator.Initialize("../../../config", "integration")

	dbURL := viper.GetString("database.url")
	db, err := gorm.Open(postgres.Open(dbURL), nil)
	s.NoError(err)
	s.dbConnection = db

	migration, err := migrate.New("file://migration", dbURL)
	s.NoError(err)
	s.dbMigration = migration

	repo := repository.New(db)
	s.useCase = usecase.New(repo)
}

func (s *useCaseTestSuite) SetupTest() {
	err := s.dbMigration.Up()
	s.NoError(err)
}

func (s *useCaseTestSuite) TearDownTest() {
	err := s.dbMigration.Down()
	s.NoError(err)
}

func (s *useCaseTestSuite) Test_CreateTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}

	err := s.useCase.CreateTodo(context.Background(), &td)
	s.NoError(err)
	s.NotZero(td.ID)
	result := model.Todo{}
	err = s.dbConnection.Where(&td).First(&result).Error
	s.NoError(err)
}

func (s *useCaseTestSuite) Test_GetTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}
	s.NoError(s.dbConnection.Create(&td).Error)

	result, err := s.useCase.GetTodoByID(context.Background(), td.ID)
	s.NoError(err)
	s.Equal(td.ID, result.ID)
	s.Equal(td.Title, result.Title)
	s.Equal(td.Description, result.Description)
}

func (s *useCaseTestSuite) Test_DeleteTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}
	s.NoError(s.dbConnection.Create(&td).Error)

	err := s.useCase.DeleteTodo(context.Background(), td.ID)
	s.NoError(err)
	result := model.Todo{}
	err = s.dbConnection.First(&result, td.ID).Error
	s.True(errors.Is(err, gorm.ErrRecordNotFound))
	s.Equal(model.Todo{}, result)
}

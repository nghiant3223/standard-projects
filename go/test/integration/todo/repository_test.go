// +build integration

package todo

import (
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nghiant3223/standard-project/internal/todo"
	"github.com/nghiant3223/standard-project/internal/todo/apperror"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/nghiant3223/standard-project/internal/todo/repository"
	"github.com/nghiant3223/standard-project/pkg/configurator"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repositoryTestSuite struct {
	suite.Suite

	repository   todo.Repository
	dbConnection *gorm.DB
	dbMigration  *migrate.Migrate
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &repositoryTestSuite{})
}

func (s *repositoryTestSuite) SetupSuite() {
	configurator.Initialize("../../../config", "integration")

	dbURL := viper.GetString("database.url")
	db, err := gorm.Open(postgres.Open(dbURL), nil)
	s.NoError(err)
	s.dbConnection = db

	migration, err := migrate.New("file://migration", dbURL)
	s.NoError(err)
	s.dbMigration = migration

	s.repository = repository.New(db)
}

func (s *repositoryTestSuite) SetupTest() {
	err := s.dbMigration.Up()
	s.NoError(err)
}

func (s *repositoryTestSuite) TearDownTest() {
	err := s.dbMigration.Down()
	s.NoError(err)
}

func (s *repositoryTestSuite) Test_CreateTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}

	err := s.repository.Create(&td)
	s.NoError(err)
	s.NotZero(td.ID)
}

func (s *repositoryTestSuite) Test_GetTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}
	s.NoError(s.repository.Create(&td))

	result, err := s.repository.Get(td.ID)
	s.NoError(err)
	s.Equal(td.ID, result.ID)
	s.Equal(td.Title, result.Title)
	s.Equal(td.Description, result.Description)
}

func (s *repositoryTestSuite) Test_DeleteTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "Must do it before mom comes home",
	}
	s.NoError(s.repository.Create(&td))

	err := s.repository.Delete(td.ID)
	s.NoError(err)
	result, err := s.repository.Get(td.ID)
	s.True(errors.Is(err, apperror.ErrNotFound))
	s.Equal(model.Todo{}, result)
}

package todo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nghiant3223/standard-project/internal/todo"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/nghiant3223/standard-project/internal/todo/repository"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repositoryTestSuite struct {
	suite.Suite

	repo    todo.Repository
	sqlMock sqlmock.Sqlmock
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &repositoryTestSuite{})
}

func (s *repositoryTestSuite) SetupSuite() {
	sqlDB, sqlMock, err := sqlmock.New()
	s.NoError(err)
	s.sqlMock = sqlMock
	config := postgres.Config{Conn: sqlDB}
	db, err := gorm.Open(postgres.New(config), nil)
	s.NoError(err)
	s.repo = repository.New(db)
}

func (s *repositoryTestSuite) Test_CreateTodo_Happy() {
	query := `INSERT INTO "todos" ("title","description") VALUES ($1,$2) RETURNING "id"`
	td := model.Todo{
		Title:       "Submit assignment",
		Description: "Assignment of CS50 will be closed at 8PM",
	}
	rows := []string{"id"}
	newID := 1
	s.sqlMock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(td.Title, td.Description).
		WillReturnRows(sqlmock.NewRows(rows).AddRow(newID))

	err := s.repo.Create(&td)
	s.NoError(err)
	s.Equal(newID, td.ID)
}

func (s *repositoryTestSuite) Test_ListTodos_Happy() {
	query := `SELECT * FROM "todos"`
	todos := []model.Todo{
		{
			ID:          183,
			Title:       "Submit assignment",
			Description: "Assignment of CS50 will be closed at 8PM",
		},
		{
			ID:          239,
			Title:       "Clean the floor",
			Description: "Do it before mom comes home",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(todos[0].ID, todos[0].Title, todos[0].Description).
		AddRow(todos[1].ID, todos[1].Title, todos[1].Description)
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	result, err := s.repo.List()
	s.NoError(err)
	s.Equal(todos, result)
}

func (s *repositoryTestSuite) Test_Get_Happy() {
	query := `SELECT * FROM "todos" WHERE "todos"."id" = $1 ORDER BY "todos"."id" LIMIT 1`
	td := model.Todo{
		ID:          138,
		Title:       "Submit assignment",
		Description: "Assignment of CS50 will be closed at 8PM",
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(td.ID, td.Title, td.Description)
	s.sqlMock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(td.ID).
		WillReturnRows(rows)

	result, err := s.repo.Get(td.ID)
	s.NoError(err)
	s.Equal(td, result)
}

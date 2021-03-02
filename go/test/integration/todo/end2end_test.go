// +build integration

package todo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nghiant3223/standard-project/internal/todo/controller"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/nghiant3223/standard-project/internal/todo/repository"
	"github.com/nghiant3223/standard-project/internal/todo/usecase"
	"github.com/nghiant3223/standard-project/pkg/configurator"
	"github.com/nghiant3223/standard-project/pkg/httpserver"
	"github.com/sebdah/goldie/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type end2endTestSuite struct {
	suite.Suite

	baseURL      string
	server       *httpserver.Server
	dbMigration  *migrate.Migrate
	dbConnection *gorm.DB
}

func TestEnd2EndTestSuite(t *testing.T) {
	suite.Run(t, &end2endTestSuite{})
}

func (s *end2endTestSuite) SetupSuite() {
	configurator.Initialize("../../../config", "integration")

	dbURL := viper.GetString("database.url")
	db, err := gorm.Open(postgres.Open(dbURL), nil)
	s.NoError(err)
	s.dbConnection = db

	migration, err := migrate.New("file://migration", dbURL)
	s.NoError(err)
	s.dbMigration = migration

	repo := repository.New(db)
	uc := usecase.New(repo)
	val := validator.New()
	ctrl := controller.New(uc, val)

	router := gin.New()
	ctrl.Register(router)
	port := viper.GetInt("port")
	s.server = httpserver.NewServer(port, router)
	s.baseURL = fmt.Sprintf("http://localhost:%d", port)

	go func() {
		err = s.server.Start(context.Background())
		s.Require().NoError(err)
	}()
}

func (s *end2endTestSuite) TearDownSuite() {
	err := s.server.Stop(context.Background())
	s.Require().NoError(err)
}

func (s *end2endTestSuite) SetupTest() {
	err := s.dbMigration.Up()
	s.NoError(err)
}

func (s *end2endTestSuite) TearDownTest() {
	err := s.dbMigration.Down()
	s.NoError(err)
}

func (s *end2endTestSuite) assertResponseBody(body io.Reader) {
	bodyData, err := ioutil.ReadAll(body)
	s.NoError(err)

	buffer := &bytes.Buffer{}
	err = json.Indent(buffer, bodyData, "", "\t")
	s.NoError(err)

	t := s.T()
	g := goldie.New(t)
	indentedBody := buffer.Bytes()
	g.Assert(t, t.Name(), indentedBody)
}

func (s *end2endTestSuite) Test_CreateTodo_Happy() {
	reqBody := `{"title":"Go to the market","description":"Buy some foods and drinks"}`
	resp, err := http.Post(s.baseURL+"/todos", gin.MIMEJSON, strings.NewReader(reqBody))
	s.NoError(err)
	defer resp.Body.Close()
	s.Equal(http.StatusOK, resp.StatusCode)
	s.assertResponseBody(resp.Body)
}

func (s *end2endTestSuite) Test_GetTodo_Happy() {
	td := model.Todo{
		Title:       "Fill the gas tank",
		Description: "Go to the gas station and fill motorcycle's gas tank",
	}
	err := s.dbConnection.Create(&td).Error
	s.NoError(err)

	url := fmt.Sprintf("%s%s%d", s.baseURL, "/todos/", td.ID)
	resp, err := http.Get(url)
	s.NoError(err)
	defer resp.Body.Close()
	s.Equal(http.StatusOK, resp.StatusCode)
	s.assertResponseBody(resp.Body)
}

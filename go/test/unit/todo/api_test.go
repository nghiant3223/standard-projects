package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/nghiant3223/standard-project/internal/todo/controller"
	mocktodo "github.com/nghiant3223/standard-project/internal/todo/gomock"
	"github.com/nghiant3223/standard-project/internal/todo/model"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/suite"
)

type apiTestSuite struct {
	suite.Suite

	ginEngine *gin.Engine
	useCase   *mocktodo.MockUseCase
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, &apiTestSuite{})
}

func (s *apiTestSuite) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())

	val := validator.New()
	s.useCase = mocktodo.NewMockUseCase(mockCtrl)
	ctrl := controller.New(s.useCase, val)

	router := gin.New()
	ctrl.Register(router)
	s.ginEngine = router
}

func (s *apiTestSuite) assertResponseBody(body io.Reader) {
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

func (s *apiTestSuite) Test_GetTodo_Happy() {
	td := model.Todo{
		ID:          1,
		Title:       "Clean the floor",
		Description: "In 30 minutes",
	}
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/todos/%d", td.ID)
	req := httptest.NewRequest("GET", url, nil)

	s.useCase.EXPECT().GetTodoByID(gomock.Any(), td.ID).Return(td, nil)
	s.ginEngine.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.assertResponseBody(w.Body)
}

func (s *apiTestSuite) Test_CreateTodo_Happy() {
	td := model.Todo{
		Title:       "Clean the floor",
		Description: "In 30 minutes",
	}
	newTD := model.Todo{
		ID:          3930,
		Title:       td.Title,
		Description: td.Description,
	}
	w := httptest.NewRecorder()
	url := "/todos"
	bodyData, err := json.Marshal(td)
	s.NoError(err)
	body := bytes.NewReader(bodyData)
	req := httptest.NewRequest("POST", url, body)

	s.useCase.EXPECT().
		CreateTodo(gomock.Any(), &td).
		SetArg(1, newTD).
		Return(nil)
	s.ginEngine.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.assertResponseBody(w.Body)
}

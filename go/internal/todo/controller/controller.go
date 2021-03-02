package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nghiant3223/standard-project/internal/todo"
	"github.com/nghiant3223/standard-project/internal/todo/apperror"
	"github.com/nghiant3223/standard-project/internal/todo/converter"
	"github.com/nghiant3223/standard-project/internal/todo/dto"
	"github.com/nghiant3223/standard-project/pkg/ginwrapper"
)

type Controller struct {
	useCase             todo.UseCase
	validator           *validator.Validate
	getTodoConverter    *converter.GetTodoConverter
	createTodoConverter *converter.CreateTodoConverter
}

func New(useCase todo.UseCase, validator *validator.Validate) *Controller {
	return &Controller{
		useCase:             useCase,
		validator:           validator,
		getTodoConverter:    converter.NewGetTodoConverter(),
		createTodoConverter: converter.NewCreateTodoTransformer(),
	}
}

func (h *Controller) Register(mux gin.IRouter) {
	mux.GET("/todos", ginwrapper.Wrap(h.listTodo))
	mux.GET("/todos/:id", ginwrapper.Wrap(h.getTodo))
	mux.POST("/todos", ginwrapper.Wrap(h.createTodo))
}

func (h *Controller) listTodo(ctx *gin.Context) *ginwrapper.Response {
	todos, err := h.useCase.ListTodos(ctx)
	if err != nil {
		return &ginwrapper.Response{Error: err}
	}
	return &ginwrapper.Response{Data: todos}
}

func (h *Controller) getTodo(ctx *gin.Context) *ginwrapper.Response {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return &ginwrapper.Response{Error: err}
	}
	td, err := h.useCase.GetTodoByID(ctx, id)
	dto := h.getTodoConverter.ToDto(td)
	return &ginwrapper.Response{Data: dto}
}

func (h *Controller) createTodo(ctx *gin.Context) *ginwrapper.Response {
	var req dto.CreateTodoRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		return &ginwrapper.Response{Error: err}
	}
	err = h.validator.Struct(req)
	if err != nil {
		return &ginwrapper.Response{Error: apperror.ErrInvalid}
	}
	td := h.createTodoConverter.ToEntity(req)
	err = h.useCase.CreateTodo(ctx, &td)
	if err != nil {
		return &ginwrapper.Response{Error: err}
	}
	return &ginwrapper.Response{Data: td}
}

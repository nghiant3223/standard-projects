// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_todo is a generated GoMock package.
package mock_todo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/nghiant3223/standard-project/internal/todo/model"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// ListTodos mocks base method
func (m *MockUseCase) ListTodos(ctx context.Context) ([]model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTodos", ctx)
	ret0, _ := ret[0].([]model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTodos indicates an expected call of ListTodos
func (mr *MockUseCaseMockRecorder) ListTodos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTodos", reflect.TypeOf((*MockUseCase)(nil).ListTodos), ctx)
}

// GetTodoByID mocks base method
func (m *MockUseCase) GetTodoByID(ctx context.Context, id int) (model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoByID", ctx, id)
	ret0, _ := ret[0].(model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoByID indicates an expected call of GetTodoByID
func (mr *MockUseCaseMockRecorder) GetTodoByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoByID", reflect.TypeOf((*MockUseCase)(nil).GetTodoByID), ctx, id)
}

// CreateTodo mocks base method
func (m *MockUseCase) CreateTodo(ctx context.Context, todo *model.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", ctx, todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTodo indicates an expected call of CreateTodo
func (mr *MockUseCaseMockRecorder) CreateTodo(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockUseCase)(nil).CreateTodo), ctx, todo)
}

// DeleteTodo mocks base method
func (m *MockUseCase) DeleteTodo(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo
func (mr *MockUseCaseMockRecorder) DeleteTodo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockUseCase)(nil).DeleteTodo), ctx, id)
}
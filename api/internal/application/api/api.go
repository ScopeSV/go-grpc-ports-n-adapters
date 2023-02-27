package api

import (
	"github.com/scopesv/todoGrpc/api/internal/application/domain"
	"github.com/scopesv/todoGrpc/api/internal/ports"
)

type Application struct {
	todo ports.TodoPort
}

func (a *Application) GetAllTodos() (*[]domain.Todo, error) {
	todo, err := a.todo.GetTodos()
	return todo, err
}

func (a *Application) SaveTodo(todo *domain.TodoRequest) error {
	return a.todo.SaveTodo(todo)
}

func (a *Application) DeleteTodo(id int64) error {
	return a.todo.DeleteTodo(id)
}

func (a *Application) SetTodoCompleted(id int64) error {
	return a.todo.SetTodoCompleted(id)
}

func NewApplication(todo ports.TodoPort) *Application {
	return &Application{
		todo: todo,
	}
}

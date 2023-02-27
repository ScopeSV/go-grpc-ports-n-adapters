package api

import (
	"github.com/scopesv/todoGrpc/server/internal/application/domain"
	"github.com/scopesv/todoGrpc/server/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func (a *Application) GetTodos() ([]*domain.Todo, error) {
	todos, err := a.db.GetAll()

	if err != nil {
		return []*domain.Todo{}, err
	}

	return todos, nil

}

func (a *Application) SaveTodo(todo *domain.Todo) (int64, error) {
	return a.db.Save(todo)
}

func (a *Application) DeleteTodo(id int64) error {
	return a.db.Delete(id)
}

func (a *Application) CompleteTodo(id int64) error {
	return a.db.SetCompleted(id)
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

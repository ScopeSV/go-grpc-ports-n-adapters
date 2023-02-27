package ports

import (
	"github.com/scopesv/todoGrpc/api/internal/application/domain"
)

type TodoPort interface {
	GetTodos() (*[]domain.Todo, error)
	SaveTodo(todo *domain.TodoRequest) error
	DeleteTodo(id int64) error
	SetTodoCompleted(id int64) error
}

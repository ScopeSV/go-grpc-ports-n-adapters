package ports

import "github.com/scopesv/todoGrpc/api/internal/application/domain"

type APIPort interface {
	GetAllTodos() (*[]domain.Todo, error)
	SaveTodo(*domain.TodoRequest) error
	DeleteTodo(int64) error
	SetTodoCompleted(int64) error
}

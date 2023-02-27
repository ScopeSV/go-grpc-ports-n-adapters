package ports

import "github.com/scopesv/todoGrpc/server/internal/application/domain"

type APIPort interface {
	GetTodos() ([]*domain.Todo, error)
	SaveTodo(todo *domain.Todo) (int64, error)
	DeleteTodo(id int64) error
	CompleteTodo(id int64) error
}

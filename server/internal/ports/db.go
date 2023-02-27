package ports

import "github.com/scopesv/todoGrpc/server/internal/application/domain"

type DBPort interface {
	GetAll() ([]*domain.Todo, error)
	Save(todo *domain.Todo) (int64, error)
	Delete(id int64) error
	SetCompleted(id int64) error
}

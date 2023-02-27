package grpc

import (
	"context"

	"github.com/scopesv/todoGrpc/server/internal/application/domain"
	"github.com/scopesv/todoGrpc/server/pkg/pb"
)

func (a *Adapter) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	newTodo := domain.NewTodo(req.Desc)
	todoID, err := a.api.SaveTodo(newTodo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTodoResponse{Id: todoID}, nil
}

func (a *Adapter) GetAllTodos(ctx context.Context, req *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	res, err := a.api.GetTodos()

	if err != nil {
		return nil, err
	}

	var todos []*pb.Todo
	for _, todo := range res {
		todos = append(todos, &pb.Todo{
			ID:        todo.ID,
			Desc:      todo.Desc,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
		})
	}

	return &pb.GetTodosResponse{Todos: todos}, nil
}

func (a *Adapter) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	err := a.api.DeleteTodo(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteTodoResponse{}, nil
}

func (a *Adapter) SetTodoCompleted(ctx context.Context, req *pb.SetTodoCompletedRequest) (*pb.SetTodoCompletedResponse, error) {
	err := a.api.CompleteTodo(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.SetTodoCompletedResponse{}, nil
}

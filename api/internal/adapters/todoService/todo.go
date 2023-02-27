package todoService

import (
	"context"

	"github.com/scopesv/todoGrpc/api/internal/application/domain"
	"github.com/scopesv/todoGrpc/api/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	todo pb.TodoServiceClient
}

func (a *Adapter) GetTodos() (*[]domain.Todo, error) {
	todos, err := a.todo.GetAllTodos(context.Background(), &pb.GetTodosRequest{})
	if err != nil {
		return nil, err
	}

	var todoList []domain.Todo
	for _, todo := range todos.Todos {
		todoList = append(todoList, domain.Todo{
			ID:        todo.ID,
			Desc:      todo.Desc,
			Completed: todo.Completed,
		})
	}

	return &todoList, nil
}

func (a *Adapter) SaveTodo(todo *domain.TodoRequest) error {
	_, err := a.todo.CreateTodo(context.Background(), &pb.CreateTodoRequest{
		Desc: todo.Desc,
	})

	return err
}

func (a *Adapter) DeleteTodo(id int64) error {
	_, err := a.todo.DeleteTodo(context.Background(), &pb.DeleteTodoRequest{
		Id: id,
	})

	return err
}

func (a *Adapter) SetTodoCompleted(id int64) error {
	_, err := a.todo.SetTodoCompleted(context.Background(), &pb.SetTodoCompletedRequest{
		Id: id,
	})

	return err
}

func NewAdapter(todoServiceUri string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(todoServiceUri, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewTodoServiceClient(conn)
	return &Adapter{todo: client}, nil
}

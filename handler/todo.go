package handler

import (
	"log"
	"time"

	pb "github.com/Aymeric-Henry/todoGrpc/proto/todo"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
	pb.UnimplementedTodoServiceServer
	Todos []*pb.TodoObject
}

func (s *Server) AddTodo(ctx context.Context, newTodo *pb.AddTodoParams) (*pb.TodoObject, error) {
	log.Printf("Received new task %s", newTodo.Task)
	todoObject := &pb.TodoObject{
		Id:        uuid.NewV1().String(),
		Task:      newTodo.Task,
		Completed: false,
		Deleted:   false,
	}
	s.Todos = append(s.Todos, todoObject)
	return todoObject, nil
}

func (s *Server) GetTodos(todo *pb.GetTodoParams, stream pb.TodoService_GetTodosServer) error {
	for _, todo := range s.Todos {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.TodoResponse{Todo: todo}); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) DeleteTodo(ctx context.Context, delTodo *pb.DeleteTodoParams) (*pb.DeleteResponse, error) {
	for _, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			todo.Deleted = true
		}
	}
	return &pb.DeleteResponse{Message: "success"}, nil
}

func (s *Server) MarkTodo(ctx context.Context, delTodo *pb.MarkTodoParams) (*pb.MarkResponse, error) {
	for _, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			todo.Completed = delTodo.Completed
		}
	}
	return &pb.MarkResponse{Message: "success"}, nil
}

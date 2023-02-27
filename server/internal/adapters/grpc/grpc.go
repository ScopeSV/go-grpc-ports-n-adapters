package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/scopesv/todoGrpc/server/config"
	"github.com/scopesv/todoGrpc/server/internal/ports"
	"github.com/scopesv/todoGrpc/server/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	pb.UnimplementedTodoServiceServer
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, a)

	if config.GetEnv() == "DEV" {
		reflection.Register(grpcServer)
	}

	log.Printf("Server is running on port %v", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: error: %v", a.port, err)
	}
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

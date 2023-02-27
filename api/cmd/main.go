package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scopesv/todoGrpc/api/config"
	"github.com/scopesv/todoGrpc/api/internal/adapters/httpServer"
	"github.com/scopesv/todoGrpc/api/internal/adapters/todoService"
	"github.com/scopesv/todoGrpc/api/internal/application/api"
)

func init() {
	godotenv.Load()
}

func main() {
	todoSrvAdapter, err := todoService.NewAdapter(config.GetTodoServiceUri())

	if err != nil {
		log.Fatalf("Failed to create connect to Todo service: %v", err)
	}

	app := api.NewApplication(todoSrvAdapter)
	httpAdapter := httpServer.NewAdapter(app, config.GetHttpPort())

	httpAdapter.StartServer()
}

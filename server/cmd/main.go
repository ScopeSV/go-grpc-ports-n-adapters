package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/scopesv/todoGrpc/server/config"
	"github.com/scopesv/todoGrpc/server/internal/adapters/db"
	"github.com/scopesv/todoGrpc/server/internal/adapters/grpc"
	"github.com/scopesv/todoGrpc/server/internal/application/api"
)

func init() {
	godotenv.Load()
}

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		log.Fatalf(err.Error())
	}

	app := api.NewApplication(dbAdapter)

	grpcAdapter := grpc.NewAdapter(app, config.GetApplicationPort())
	grpcAdapter.Run()
}

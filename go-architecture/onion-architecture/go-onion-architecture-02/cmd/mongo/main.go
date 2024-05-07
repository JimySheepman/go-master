package main

import (
	"context"
	"log"

	"golang_startup/internal/common"
	"golang_startup/internal/config"
	"golang_startup/internal/controller"
	"golang_startup/internal/database/mongodb"
	"golang_startup/internal/server"
	"golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := mongodb.New(context.Background(), cfg.Database.Address, cfg.Database.Name)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	timer := common.NewTimer()
	services := service.New(database, timer)

	controllers := controller.New(services)

	srv := server.New(controllers, cfg.GetServerAddress())

	srv.ServeHTTP()
}

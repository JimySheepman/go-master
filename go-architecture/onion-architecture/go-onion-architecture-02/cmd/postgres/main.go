package main

import (
	"log"

	"golang_startup/internal/common"
	"golang_startup/internal/config"
	"golang_startup/internal/controller"
	"golang_startup/internal/database/postgres"
	"golang_startup/internal/server"
	"golang_startup/internal/service"
)

func main() {
	cfg := config.Load()

	database, err := postgres.New(cfg.Database.Address)
	if err != nil {
		log.Fatalf("failed to initialise the database client: %v", err)
	}

	err = database.Migrate(cfg.Database.Name, cfg.Database.MigrationsPath)
	if err != nil {
		log.Fatalf("failed initialise database migrations: %v", err)
	}

	timer := common.NewTimer()
	services := service.New(database, timer)

	controllers := controller.New(services)

	srv := server.New(controllers, cfg.GetServerAddress())

	srv.ServeHTTP()
}

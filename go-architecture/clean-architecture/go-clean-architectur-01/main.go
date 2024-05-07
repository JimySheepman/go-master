package main

import (
	"os"
	"time"

	"go-clean-architecture/infrastructure"
	"go-clean-architecture/infrastructure/database"
	"go-clean-architecture/infrastructure/log"
	"go-clean-architecture/infrastructure/router"
	"go-clean-architecture/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}

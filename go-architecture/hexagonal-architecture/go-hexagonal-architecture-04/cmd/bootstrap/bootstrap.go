package bootstrap

import (
	"context"
	"fmt"
	database "go-hexagonal/internal/adapters/driven"
	"go-hexagonal/internal/core/application"
	"go-hexagonal/internal/platform/server"
	mysqldb "go-hexagonal/internal/platform/storage/mysql"
	"go-hexagonal/pkg/config"
	"log"
)

func Run() error {

	err := config.LoadConfig()
	if err != nil {
		return err
	}
	fmt.Println("Web server ready!")

	ctx := context.Background()
	db, err := config.ConfigDb(ctx)

	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	userRepository := mysqldb.NewUserRepository(db, config.Cfg.DbTimeout)
	userAdapter := database.NewUserAdapter(userRepository)
	userService := application.NewUserService(userAdapter)

	ctx, srv := server.NewServer(context.Background(), config.Cfg.Host, config.Cfg.Port, config.Cfg.ShutdownTimeout, server.AppService{
		UserService: userService,
	})

	return srv.Run(ctx)
}

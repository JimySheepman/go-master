package main

import (
	"go-clean-arch-v2/config"
	"go-clean-arch-v2/database"
	"go-clean-arch-v2/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}

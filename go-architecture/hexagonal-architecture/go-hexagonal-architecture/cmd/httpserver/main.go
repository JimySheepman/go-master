package main

import (
	"go-hexagonal-architecture/internal/core/services/gamesrv"
	"go-hexagonal-architecture/internal/handlers/gamehdl"
	"go-hexagonal-architecture/internal/repositories/gamesrepo"
	"go-hexagonal-architecture/pkg/uidgen"

	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}

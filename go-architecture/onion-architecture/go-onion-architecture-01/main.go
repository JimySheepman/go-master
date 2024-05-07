package main

import (
	db "go-onion-architecture-sample/infrastructure"
	"go-onion-architecture-sample/registory"
	ui "go-onion-architecture-sample/userinterface"
)

func main() {
	// db
	db := db.NewDBClient("./go_onion_architecture.db")
	defer db.Client.Close()

	// registory
	registory := registory.NewRegistory(db)

	// echo
	apiClient := ui.NewApiClient(registory)
	apiClient.RegisterRoute()
	apiClient.Start()
}

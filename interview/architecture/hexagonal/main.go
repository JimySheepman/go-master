package main

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/hexagonal/adapters/api"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/hexagonal/adapters/repository"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/hexagonal/application"
	"log"
	"net/http"
)

func main() {
	userRepo := repository.NewInMemoryUserRepository()
	userService := application.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.CreateUser(w, r)
		} else if r.Method == http.MethodGet {
			userHandler.GetUserByID(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

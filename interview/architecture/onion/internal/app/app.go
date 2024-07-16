package app

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/repository"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/service"
	https "github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/transport/http"
	"net/http"
)

// StartApplication başlatır
func StartApplication() {
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := https.NewUserHandler(userService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.CreateUser(w, r)
		} else if r.Method == http.MethodGet {
			userHandler.GetUserByID(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

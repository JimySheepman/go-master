package app

import (
	https "github.com/JimySheepman/go-master/go-algorithm/architecture/clean/internal/delivery/http"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/clean/internal/repository"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/clean/internal/usecase"
	"net/http"
)

// StartApplication başlatır
func StartApplication() {
	userRepo := repository.NewInMemoryUserRepository()
	userUsecase := usecase.NewUserUseCase(userRepo)
	userHandler := https.NewUserHandler(userUsecase)

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

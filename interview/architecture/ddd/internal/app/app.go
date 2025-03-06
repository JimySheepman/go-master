package app

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/application"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/infrastructure/persistence"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/infrastructure/web"
	"net/http"
)

// StartApplication başlatır
func StartApplication() {
	orderRepo := persistence.NewMemoryOrderRepository()
	orderService := application.NewOrderService(orderRepo)
	orderHandler := web.NewOrderHandler(orderService)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			orderHandler.CreateOrder(w, r)
		} else if r.Method == http.MethodGet {
			orderHandler.GetOrderByID(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

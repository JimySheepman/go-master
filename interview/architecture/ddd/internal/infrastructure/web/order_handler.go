package web

import (
	"encoding/json"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/application"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/domain"
	"net/http"
)

// OrderHandler, HTTP isteklerini yöneten handler
type OrderHandler struct {
	orderService *application.OrderService
}

// NewOrderHandler, yeni bir OrderHandler oluşturur
func NewOrderHandler(orderService *application.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// CreateOrder, HTTP POST isteğini işler
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CustomerID string  `json:"customer_id"`
		Amount     float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := h.orderService.CreateOrder(req.CustomerID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrderByID, HTTP GET isteğini işler
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		if err == domain.ErrOrderNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(order)
}

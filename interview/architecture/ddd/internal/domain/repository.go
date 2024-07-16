package domain

// OrderRepository interface
type OrderRepository interface {
	Save(order Order) error
	FindByID(id string) (Order, error)
}

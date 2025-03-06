// Package request holds request model structs and methods.
package request

// UserUpdate - request model for a user update
// swagger:model UserUpdate
type UserUpdate struct {
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

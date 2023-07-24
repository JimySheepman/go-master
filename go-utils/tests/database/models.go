package database

import (
	"log"
	"time"
)

type TestModel struct {
	ID            int64
	TransactionID int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Description   string
	Amount        float64
}

type testModelRepo struct {
}

func NewTestModelRepo() *testModelRepo {
	return &testModelRepo{}
}

func (r *testModelRepo) Create(t *TestModel) error {
	log.Println("model is created")
	return nil
}

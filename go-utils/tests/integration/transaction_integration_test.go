package integration

import (
	"testing"
	"time"

	"github.com/JimySheepman/go-master/go-utils/tests/database"
	"github.com/stretchr/testify/require"
)

var TestBeneficiaryInn = "test"

func Test_Integration_TransactionImpl_Create(t *testing.T) {
	repo := database.NewTestModelRepo()

	ta := &database.TestModel{
		TransactionID: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Description:   "Test",
		Amount:        1,
	}

	err := repo.Create(ta)
	require.Nil(t, err)
}

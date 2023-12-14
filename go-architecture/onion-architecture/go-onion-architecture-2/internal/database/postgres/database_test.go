//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

const (
	migrationsPath = "file://../../../migrations/postgres"
	postgresqlURL  = "postgres://docker:example@%s/test_db?sslmode=disable"

	postgresDB   = "test_db"
	postgresUser = "docker"
	postgresPass = "example"
)

func setup(t *testing.T) (testcontainers.Container, error) {
	t.Helper()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14-alpine3.15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       postgresDB,
			"POSTGRES_USER":     postgresUser,
			"POSTGRES_PASSWORD": postgresPass,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func shutdown(t *testing.T, container testcontainers.Container) {
	t.Helper()

	err := container.Terminate(context.Background())
	if err != nil {
		t.Logf("error tearing down container: %v", err)
	}
}

func buildClient(t *testing.T, container testcontainers.Container) *Database {
	t.Helper()

	endpoint, err := container.Endpoint(context.Background(), "")
	require.NoError(t, err)

	dbUrl := fmt.Sprintf(postgresqlURL, endpoint)

	db, err := New(dbUrl)
	require.NoError(t, err)

	err = db.Migrate("users", migrationsPath)
	require.NoError(t, err)

	return db
}

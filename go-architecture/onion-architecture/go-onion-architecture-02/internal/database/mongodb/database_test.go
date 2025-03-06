//go:build integration
// +build integration

package mongodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	mongoUsername     = "docker"
	mongoPassword     = "example"
	mongoRootName     = "root"
	mongoRootPassword = "example"
	mongoAuthSource   = "admin"

	databaseTest = "dbTest"
)

func setup(t *testing.T) (testcontainers.Container, error) {
	t.Helper()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.2.11",
		ExposedPorts: []string{"27017/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": mongoUsername,
			"MONGO_INITDB_ROOT_PASSWORD": mongoPassword,
			"MONGO_USER_ROOT_NAME":       mongoRootName,
			"MONGO_USER_ROOT_PASSWORD":   mongoRootPassword,
			"MONGO_AUTH_SOURCE":          mongoAuthSource,
		},
		WaitingFor: wait.ForListeningPort("27017/tcp"),
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

	dbUrl := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin&ssl=false",
		mongoUsername, mongoPassword, endpoint, databaseTest)

	db, err := New(context.Background(), dbUrl, databaseTest)
	require.NoError(t, err)

	return db
}

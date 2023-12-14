//go:build integration
// +build integration

package mysql

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	migrationsPath = "file://../../../migrations/mysql"
	mysqlURL       = "docker:example@tcp(%s)/test_db"

	mysqlDB   = "test_db"
	mysqlUser = "docker"
	mysqlPass = "example"
)

func setup(t *testing.T) (testcontainers.Container, error) {
	t.Helper()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:5.7",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_DATABASE":      mysqlDB,
			"MYSQL_USER":          mysqlUser,
			"MYSQL_PASSWORD":      mysqlPass,
			"MYSQL_ROOT_PASSWORD": mysqlPass,
		},
		WaitingFor: wait.ForListeningPort("3306/tcp"),
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

	dbUrl := fmt.Sprintf(mysqlURL, endpoint)

	log.Println(dbUrl)

	db, err := New(dbUrl)
	require.NoError(t, err)

	err = db.Migrate("users", migrationsPath)
	require.NoError(t, err)

	return db
}

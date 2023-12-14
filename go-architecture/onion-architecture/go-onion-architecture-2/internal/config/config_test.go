package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	databaseAddressTest    = "mysql://localhost:3306/dbname"
	databaseNameTest       = "dbname"
	databaseMigrationsPath = "file://migrations/postgres"

	serverHostTest = "localhost"
	serverPortTest = "8080"

	serverAddressTest = "localhost:8080"
)

var expectedCfg = Config{
	Database: Database{
		Address:        databaseAddressTest,
		Name:           databaseNameTest,
		MigrationsPath: databaseMigrationsPath,
	},
	Server: Server{
		Host: serverHostTest,
		Port: serverPortTest,
	},
}

func TestLoad(t *testing.T) {
	// Test Database fields
	err := os.Setenv(databaseAddressEnv, databaseAddressTest)
	require.NoError(t, err)

	err = os.Setenv(databaseNameEnv, databaseNameTest)
	require.NoError(t, err)

	err = os.Setenv(migrationsPathEnv, databaseMigrationsPath)
	require.NoError(t, err)

	// Test server fields
	err = os.Setenv(serverHostEnv, serverHostTest)
	require.NoError(t, err)

	err = os.Setenv(serverPortEnv, serverPortTest)
	require.NoError(t, err)

	// Load configs
	cfg := Load()

	assert.Equal(t, expectedCfg, cfg)

	assert.Equal(t, serverAddressTest, cfg.GetServerAddress())
}

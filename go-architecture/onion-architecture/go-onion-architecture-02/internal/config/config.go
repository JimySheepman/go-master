// Package config holds configuration global data to be initialized.
package config

import (
	"fmt"
	"os"
)

const (
	serverHostEnv = "HOST"
	serverPortEnv = "PORT"

	databaseAddressEnv = "DB_ADDRESS"
	databaseNameEnv    = "DB_NAME"

	migrationsPathEnv = "MIGRATIONS_PATH"
)

// Database - database related config data
type Database struct {
	Address        string
	Name           string
	MigrationsPath string
}

// Server - server related config data
type Server struct {
	Host string
	Port string
}

// Config - holds global configs
type Config struct {
	Database Database
	Server   Server
}

// Load - loads configurations data
func Load() Config {
	return Config{
		Database: Database{
			Address:        os.Getenv(databaseAddressEnv),
			Name:           os.Getenv(databaseNameEnv),
			MigrationsPath: os.Getenv(migrationsPathEnv),
		},
		Server: Server{
			Host: os.Getenv(serverHostEnv),
			Port: os.Getenv(serverPortEnv),
		},
	}
}

// GetServerAddress - returns server address based on Server configs
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

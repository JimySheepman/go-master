// Package postgres holds PostgreSQL database and repository methods.
package postgres

import (
	"database/sql"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	postgresMigration "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Migrate source import
	_ "github.com/lib/pq"                                // PostgreSQL driver
)

const (
	postgresDriveName = "postgres"

	migrationsTable = "migrations"
	schemaName      = "public"
)

// Database - PostgreSQL database struct
type Database struct {
	client *sql.DB
}

// New creates a new PostgreSQL database struct
func New(address string) (*Database, error) {
	db, err := sql.Open(postgresDriveName, address)
	if err != nil {
		return &Database{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Database{}, fmt.Errorf("error pinging database connection: %s", err.Error())
	}

	return &Database{client: db}, nil
}

// Migrate - Runs a database migration for a given database and migrations path with the SQL files
// It returns an error in case of failure
func (db *Database) Migrate(databaseName string, migrationsPath string) error {
	postgresConfig := postgresMigration.Config{
		SchemaName:      schemaName,
		MigrationsTable: migrationsTable,
		DatabaseName:    databaseName,
	}

	driver, err := postgresMigration.WithInstance(db.client, &postgresConfig)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		return nil
	}

	return err
}

// Package mysql holds MySQL database and repository methods.
package mysql

import (
	"database/sql"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	mysqlMigration "github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/go-sql-driver/mysql"                   // MySQL driver
	_ "github.com/golang-migrate/migrate/v4/source/file" // Migrate source import
)

const (
	mysqlDriveName  = "mysql"
	migrationsTable = "migrations"
)

// Database - MySQL database struct
type Database struct {
	client *sql.DB
}

// New creates a new MySQL database struct
func New(address string) (*Database, error) {
	db, err := sql.Open(mysqlDriveName, address)
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
	postgresConfig := mysqlMigration.Config{
		MigrationsTable: migrationsTable,
		DatabaseName:    databaseName,
	}

	driver, err := mysqlMigration.WithInstance(db.client, &postgresConfig)
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

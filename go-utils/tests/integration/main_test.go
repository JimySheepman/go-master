package integration

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/JimySheepman/go-master/go-utils/tests/config"
	"github.com/JimySheepman/go-master/go-utils/tests/database"
)

const (
	createDBQuery = "CREATE DATABASE [master.test]"
	driverName    = "sqlserver"
	driver        = "mssql"
	ipAddr        = "127.0.0.1"
	dbUsername    = "sa"
	password      = "test.123456"
	dbName        = "master.test"
)

var migrationPaths = []string{
	"../migrations/000001.up.sql",
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	c, err := CreateMsSqlContainer(ctx, ContainerRequest{
		Password: password,
		DbName:   "[master.test]",
	})
	if err != nil {
		log.Fatal("crate container: ", err)
	}

	cs, err := c.GoConnectionString(ctx)
	if err != nil {
		log.Fatal("get connection string: ", err)
	}

	db, err := sql.Open(driverName, cs)
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer db.Close()

	_, err = config.Load()
	if err != nil {
		os.Exit(0)
	}

	_, e := db.Exec(createDBQuery)
	if e != nil {
		log.Fatal("exec: ", e)
	}
	_, p, _ := c.HostAndPort(ctx)

	conn, err := database.InitDBTest(driver, ipAddr, p.Port(), dbUsername, password, dbName, cs)
	if err != nil {
		os.Exit(0)
	}

	for _, path := range migrationPaths {
		ReadFileAndExec(path, conn, true)
	}

	os.Exit(m.Run())
}

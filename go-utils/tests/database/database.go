package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/birlesikodeme/gorp/v3"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

var db gorpClient

type gorpClient interface {
	gorp.SqlExecutor
	AddTableWithName(i interface{}, name string) *gorp.TableMap
	Begin() (*gorp.Transaction, error)
	TraceOn(prefix string, logger gorp.GorpLogger)
}

func InitDBTest(driver string, ipAddr string, portNo string, username string, password string, dbName string, address string) (*sql.DB, error) {
	var (
		ip   = ipAddr
		port = portNo
		user = username
		pass = password
		name = dbName
		err  error
		conn *sql.DB
	)

	conn, err = sql.Open("mssql", fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s", ip, port, user, pass, name))
	if err != nil {
		return nil, fmt.Errorf("database connection error: %s", err.Error())
	}

	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("database connection error: %s", err.Error())
	}

	db = &gorp.DbMap{Db: conn, Dialect: gorp.SqlServerDialect{}, ExpandSliceArgs: true}
	db.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	return conn, nil
}

package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"auth/infra/rdb"

	"github.com/google/wire"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(mysql:3306)/shellingford?parseTime=true")
	if err != nil {
		panic(err)
	}
}

var Set = wire.NewSet(
	rdb.NewUserRepository,
	rdb.NewAccountRepository,
	rdb.NewSessionRepository,
	rdb.NewUserQueryService,
	rdb.NewSessionQueryService,
	wire.Value(DB),
)

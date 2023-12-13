package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

var Cfg config

type config struct {
	DbUser          string        `default:"enviadoc"`
	DbPass          string        `default:"enviadoc"`
	DbHost          string        `default:"0.0.0.0"`
	DbPort          string        `default:"3306"`
	DbName          string        `default:"enviadoc"`
	DbTimeout       time.Duration `default:"10s"`
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
}

func LoadConfig() error {
	if err := envconfig.Process("IRIS", &Cfg); err != nil {
		return err
	}

	return nil
}

func ConfigDb(ctx context.Context) (*sql.DB, error) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", Cfg.DbUser, Cfg.DbPass, Cfg.DbHost, Cfg.DbPort, Cfg.DbName)
	log.Println("uri: ", mysqlURI)

	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Println("Failed database connection")
		panic(err)
	}

	log.Println("Successfully Connected to MySQL database")

	db.SetConnMaxLifetime(time.Minute * 4)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

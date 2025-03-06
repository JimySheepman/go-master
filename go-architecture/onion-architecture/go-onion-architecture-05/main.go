package main

import (
	"flag"
	"fmt"

	"go-echo-rest-sample/conf"
	"go-echo-rest-sample/interactor"
	"go-echo-rest-sample/presenter/http/middleware"
	"go-echo-rest-sample/presenter/http/router"

	_ "github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var runServer = flag.Bool("server", false, "production is -server option require")

func main() {
	flag.Parse()
	conf.NewConfig(*runServer)

	e := echo.New()
	conn := conf.NewDBConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}

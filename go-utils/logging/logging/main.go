package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/JimySheepman/go-master/go-utils/logging/controllers"
	"github.com/JimySheepman/go-master/go-utils/logging/database"
	"github.com/JimySheepman/go-master/go-utils/logging/pkg/log"
	"github.com/JimySheepman/go-master/go-utils/logging/pkg/pkg1"
	"github.com/JimySheepman/go-master/go-utils/logging/services"
	"go.uber.org/zap"
)

func main() {
	logger := createLogger()
	defer logger.Sync()

	repo := database.NewDummyDatabase()
	ctrl := controllers.NewDummyControllers(repo, logger)
	srv := services.NewDummyService(ctrl)

	transferHandler := func(w http.ResponseWriter, req *http.Request) {
		for i := 0; i < 5; i++ {
			err := srv.CreateTransfer(i)
			json.NewEncoder(w).Encode(err)
		}
	}

	canceledHandler := func(w http.ResponseWriter, req *http.Request) {
		for i := 0; i < 5; i++ {
			err := srv.CanceledTransfer(i)
			logger.Info("CanceledTransfer", zap.Any("err", err))
			json.NewEncoder(w).Encode(err)
		}
	}

	errorCasting()

	http.HandleFunc("/transfer", transferHandler)
	http.HandleFunc("/canceled", canceledHandler)

	logger.Sugar().Fatal(http.ListenAndServe(":3000", nil))
}

func createLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	zap.NewProductionConfig()
	return zap.Must(config.Build())
}

func fileLogger() {
	file, err := os.OpenFile("./demo1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, log.InfoLevel)
	log.ResetDefault(logger)
	defer log.Sync()
	log.Info("demo1:", log.String("app", "start ok"),
		log.Int("major version", 2))
	pkg1.Foo()
}

func demoLogger() {
	log.Info("demo1:",
		log.String("app", "start ok"),
		log.Int("major version", 2),
	)
}

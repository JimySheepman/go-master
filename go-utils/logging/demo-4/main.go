package main

import (
	"context"
	"fmt"
	"test/controllers"
	"test/databases"
	"test/pkg/anotherlogger"
	"test/pkg/extractor"
	"test/services"
	"time"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap/zapcore"
)

const (
	appName = "salary-api"
)

type config struct {
	Title           string
	EnvironmentType string `env:"ENVIRONMENT_TYPE" env-description:"application environment type" env-default:"development"`
	Port            int    `env:"PORT" env-description:"application RPC port address" env-default:"8686"`
	RestPort        int    `env:"REST_PORT" env-description:"application REST server port address" env-default:"8787"`
	LogLevel        int    `env:"LOG_LEVEL" env-description:"log level" env-default:"-1"`
	LogEncoding     string `env:"LOG_ENCODING" env-description:"log encoding" env-default:"console"`
	GraylogAddr     string `env:"GRAYLOG_ADDR" env-description:"Graylog addr:port" env-default:"127.0.0.1:12201"`
	DBHost          string `env:"DB_HOST" env-description:"database host address" env-default:"localhost"`
	DBName          string `env:"DB_NAME" env-description:"database schema name" env-default:"master"`
	DBPort          int    `env:"DB_PORT" env-description:"database port" env-default:"1433"`
	DBUsername      string `env:"DB_USERNAME" env-description:"database username" env-default:"sa"`
	DBPassword      string `env:"DB_PASSWORD" env-description:"database password" env-default:"Mysql!Server02"`
}

func main() {
	ctx := grpc_ctxtags.SetInContext(
		context.Background(),
		grpc_ctxtags.NewTags().
			Set("ip", string("127.0.0.1")).
			Set("requestID", string("abcdefghijklmn")),
	)

	ip, _ := extractor.IpExtractor(ctx)
	rid, _ := extractor.RequestIDExtractor(ctx)

	log := anotherlogger.Get().
		With(
			zapcore.Field{
				Key:       "requesdId",
				Type:      zapcore.StringType,
				Integer:   0,
				String:    rid,
				Interface: nil,
			},
		)

	repo := databases.NewDatabase()
	ctrl := controllers.NewController(repo)
	srv := services.NewService(ctrl)

	ctx = anotherlogger.WithCtx(ctx, log)

	start := time.Now()
	err := srv.DoesAnythingService(ctx)
	time.Sleep(3750 * time.Millisecond)
	latency := time.Since(start)

	log.Sugar().Errorw(
		fmt.Sprintf("/services/DoesAnythingService took %v", latency),
		"fullMethod", "/services/DoesAnythingService",
		"ip", ip,
		"latency", latency,
		"requestID", rid,
		"request", map[string]string{"name": "tester-name"},
		"error", err,
	)
}

/*
func getLogger(c *config) logger.Logger {
	var loggerOpts []log.LoggerOption

	loggerOpts = append(loggerOpts, log.WithIO(os.Stdout, c.LogLevel, c.EnvironmentType, c.LogEncoding))

	if c.GraylogAddr != "" {
		loggerOpts = append(loggerOpts, log.WithGraylogViaTCP(c.LogLevel, c.GraylogAddr))
	}

	log.InitLogger(appName, loggerOpts...)

	return log.GetLogger()
}
*/

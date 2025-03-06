package anotherlogger

import (
	"context"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ctxKey struct{}

var once sync.Once

var logger *zap.Logger

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Get() *zap.Logger {
	once.Do(func() {

		lumberjackLogger := &lumberjack.Logger{
			Filename:   "logs/log.json",
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		}

		go func() {
			for {
				//<-time.After(time.Hour * 24)
				<-time.After(time.Second * 3)
				lumberjackLogger.Rotate()
			}
		}()

		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		//productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		//developmentCfg := zap.NewDevelopmentEncoderConfig()
		//developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		core := zapcore.NewTee(
			zapcore.NewCore(
				//zapcore.NewConsoleEncoder(developmentCfg),
				zapcore.NewJSONEncoder(productionCfg),
				zapcore.AddSync(os.Stdout),
				zap.NewAtomicLevelAt(zap.DebugLevel),
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(productionCfg),
				zapcore.AddSync(lumberjackLogger),
				zap.NewAtomicLevelAt(zap.DebugLevel),
			).With(
				[]zapcore.Field{
					zap.Int("pid", os.Getpid()),
				},
			),
		)

		logger = zap.New(core, zap.AddCaller())
	})

	return logger
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}

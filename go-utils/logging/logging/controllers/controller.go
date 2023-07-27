package controllers

import (
	"errors"
	"fmt"

	"github.com/JimySheepman/go-master/go-utils/logging/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type repoClient interface {
	Create() error
}

type loggerClient interface {
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Core() zapcore.Core
	DPanic(msg string, fields ...zapcore.Field)
	Debug(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Level() zapcore.Level
	Log(lvl zapcore.Level, msg string, fields ...zapcore.Field)
	Named(s string) *zap.Logger
	Panic(msg string, fields ...zapcore.Field)
	Sugar() *zap.SugaredLogger
	Sync() error
	Warn(msg string, fields ...zapcore.Field)
	With(fields ...zapcore.Field) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
}

type dummyControllers struct {
	repoClient repoClient
	logger     loggerClient
}

func NewDummyControllers(repoClient repoClient, loggerClient loggerClient) *dummyControllers {
	return &dummyControllers{
		repoClient: repoClient,
		logger:     loggerClient,
	}
}

func (c *dummyControllers) CreateTransfer(code int) error {
	switch code {
	case 1:
		err := api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": fmt.Errorf("error : %w", errors.New("test error")).Error(),
		})
		c.logger.Info("CreateTransfer", zap.Any("err", err))
		return err
	case 2:
		err := api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": fmt.Errorf("error : %w", errors.New("test error")).Error(),
		})
		c.logger.Info("CreateTransfer", zap.Any("err", err))
		return err
	case 3:
		err := api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": fmt.Errorf("error : %w", errors.New("test error")).Error(),
		})
		c.logger.Info("CreateTransfer", zap.Any("err", err))
		return err
	case 4:
		err := api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": fmt.Errorf("error : %w", errors.New("test error")).Error(),
		})
		c.logger.Info("CreateTransfer", zap.Any("err", err))
		return err
	default:
		return fmt.Errorf("unknown")
	}
}

func (c *dummyControllers) CanceledTransfer(code int) error {
	switch code {
	case 1:
		return api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": "test error",
		})
	case 2:
		return api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": "test error",
		})
	case 3:
		return api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": "test error",
		})
	case 4:
		return api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
			"error": "test error",
		})
	default:
		return fmt.Errorf("unknown")
	}
}

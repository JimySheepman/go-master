package wrapper

import (
	"context"
	"test/pkg/extractor"
	"test/pkg/logger"
)

type wrapperLogClient struct {
}

func NewWrapperLogClient() *wrapperLogClient {
	return &wrapperLogClient{}
}

func (c *wrapperLogClient) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	pid, _ := extractor.RequestIDExtractor(ctx)
	keysAndValues = append(keysAndValues, "requestID", pid)
	logger.GetLogger2().Infow(
		msg,
		keysAndValues...,
	)
}

func (c *wrapperLogClient) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	pid, _ := extractor.RequestIDExtractor(ctx)
	logger.GetLogger2().Errorw(
		msg,
		"requestID", pid,
		keysAndValues,
	)
}

// TODO: interceptorden logger oluştur aşağıya ver request idnin taşınıp taşınmadığını göre
// TODO: context teki fieldları logluyor mudur?

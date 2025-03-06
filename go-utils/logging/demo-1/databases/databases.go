package databases

import (
	"context"
	"fmt"
	"test/pkg/logger"
	"time"
)

type database struct {
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) DoesAnythingDatabase(ctx context.Context) error {
	start := time.Now()
	defer func() {
		logger.GetLogger2().Infow(
			"DoesAnythingDatabase error",
			"latency", time.Since(start),
		)
	}()

	logger.GetLogger2().Errorw(
		"DoesAnythingDatabase error",
	)

	time.Sleep(1 * time.Second)
	return fmt.Errorf("DoesAnythingDatabase error")
}

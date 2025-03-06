package databases

import (
	"context"
	"fmt"
	"time"
)

type logClient interface {
	Errorw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type database struct {
	log logClient
}

func NewDatabase(log logClient) *database {
	return &database{
		log: log,
	}
}

func (d *database) DoesAnythingDatabase(ctx context.Context) error {
	start := time.Now()
	defer func() {
		d.log.Infow("DoesAnythingDatabase error", "latency", time.Since(start))
	}()

	d.log.Errorw("DoesAnythingDatabase error")

	time.Sleep(1 * time.Second)
	return fmt.Errorf("DoesAnythingDatabase error")
}

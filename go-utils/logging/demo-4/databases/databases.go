package databases

import (
	"context"
	"fmt"
	"test/pkg/anotherlogger"
	"time"
)

type database struct {
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) DoesAnythingDatabase(ctx context.Context) error {
	l := anotherlogger.FromCtx(ctx).Sugar()

	start := time.Now()
	defer func() {
		l.Infow("DoesAnythingDatabase error", "latency", time.Since(start))
	}()

	l.Errorw("DoesAnythingDatabase error")

	time.Sleep(1 * time.Second)
	return fmt.Errorf("DoesAnythingDatabase error")
}

package controllers

import (
	"context"
	"fmt"
)

type logClient interface {
	Errorw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type databaseClient interface {
	DoesAnythingDatabase(ctx context.Context) error
}

type controller struct {
	databaseClient databaseClient
	log            logClient
}

func NewController(databaseClient databaseClient, log logClient) *controller {
	return &controller{
		databaseClient: databaseClient,
		log:            log,
	}
}

func (c *controller) DoesAnythingController(ctx context.Context) error {
	if err := c.databaseClient.DoesAnythingDatabase(ctx); err != nil {
		c.log.Errorw(
			fmt.Sprintf("databaseClient.DoesAnythingDatabase error: %s", err),
		)

		return fmt.Errorf("DoesAnythingController error: %w", err)
	}

	return nil
}

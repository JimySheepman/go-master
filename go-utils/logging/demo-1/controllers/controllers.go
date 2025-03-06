package controllers

import (
	"context"
	"fmt"
	"test/pkg/logger"
)

type databaseClient interface {
	DoesAnythingDatabase(ctx context.Context) error
}

type controller struct {
	databaseClient databaseClient
}

func NewController(databaseClient databaseClient) *controller {
	return &controller{
		databaseClient: databaseClient,
	}
}

func (c *controller) DoesAnythingController(ctx context.Context) error {
	if err := c.databaseClient.DoesAnythingDatabase(ctx); err != nil {
		logger.GetLogger2().Errorw(
			fmt.Sprintf("databaseClient.DoesAnythingDatabase error: %s", err),
		)

		return fmt.Errorf("DoesAnythingController error: %w", err)
	}

	return nil
}

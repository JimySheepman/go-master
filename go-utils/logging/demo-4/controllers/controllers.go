package controllers

import (
	"context"
	"fmt"
	"test/pkg/anotherlogger"
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
	l := anotherlogger.FromCtx(ctx).Sugar()

	if err := c.databaseClient.DoesAnythingDatabase(ctx); err != nil {
		l.Errorw(
			fmt.Sprintf("databaseClient.DoesAnythingDatabase error: %s", err),
		)

		return fmt.Errorf("DoesAnythingController error: %w", err)
	}

	return nil
}

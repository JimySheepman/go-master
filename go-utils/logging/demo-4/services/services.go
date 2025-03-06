package services

import (
	"context"
	"fmt"
	"test/pkg/anotherlogger"
)

type controllerClient interface {
	DoesAnythingController(ctx context.Context) error
}

type service struct {
	controllerClient controllerClient
}

func NewService(controllerClient controllerClient) *service {
	return &service{
		controllerClient: controllerClient,
	}
}

func (s *service) DoesAnythingService(ctx context.Context) error {
	l := anotherlogger.FromCtx(ctx).Sugar()

	if err := s.controllerClient.DoesAnythingController(ctx); err != nil {
		l.Errorw(fmt.Sprintf("DoesAnythingService error: %s", err))
		return fmt.Errorf("DoesAnythingService error: %w", err)
	}

	return nil
}

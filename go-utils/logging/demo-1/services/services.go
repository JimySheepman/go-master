package services

import (
	"context"
	"fmt"
	"test/pkg/logger"
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
	if err := s.controllerClient.DoesAnythingController(ctx); err != nil {
		logger.GetLogger2().Errorw(
			fmt.Sprintf("DoesAnythingService error: %s", err),
		)

		return fmt.Errorf("DoesAnythingService error: %w", err)
	}

	return nil
}

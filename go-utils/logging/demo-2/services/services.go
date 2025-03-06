package services

import (
	"context"
	"fmt"
)

type logClient interface {
	Errorw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
}

type controllerClient interface {
	DoesAnythingController(ctx context.Context) error
}

type service struct {
	controllerClient controllerClient
	log              logClient
}

func NewService(controllerClient controllerClient, log logClient) *service {
	return &service{
		controllerClient: controllerClient,
		log:              log,
	}
}

func (s *service) DoesAnythingService(ctx context.Context) error {
	if err := s.controllerClient.DoesAnythingController(ctx); err != nil {
		s.log.Errorw(fmt.Sprintf("DoesAnythingService error: %s", err))
		return fmt.Errorf("DoesAnythingService error: %w", err)
	}

	return nil
}

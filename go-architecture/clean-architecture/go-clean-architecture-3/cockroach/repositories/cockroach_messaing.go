package repositories

import (
	"go-clean-arch-v2/cockroach/entities"

	"github.com/labstack/gommon/log"
)

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}

type cockroachFCMMessaging struct{}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockroachFCMMessaging{}
}

func (c *cockroachFCMMessaging) PushNotification(m *entities.CockroachPushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Debugf("Pushed FCM notification with data: %v", m)

	return nil
}

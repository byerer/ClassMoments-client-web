package notification

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/service/notification"
)

type notificationRepo struct {
	data *data.Data
}

func NewNotificationRepo(
	data *data.Data,
) notification.NotificationRepo {
	return &notificationRepo{
		data: data,
	}
}

func (nr *notificationRepo) AddNotification(notification *entity.Notification) error {
	result := nr.data.DB.Create(notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

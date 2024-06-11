package notification

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	"time"
)

type NotificationRepo interface {
	AddNotification(notification *entity.Notification) error
}

type NotificationService interface {
	AddNotification(notificationReq *schema.NotificationReq) error
}

type notificationService struct {
	notificationRepo NotificationRepo
}

func NewNotificationService(
	notificationRepo NotificationRepo,
) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

func (ns *notificationService) AddNotification(notificationReq *schema.NotificationReq) error {
	notification := &entity.Notification{
		UserID:   notificationReq.UserID,
		MomentID: notificationReq.MomentID,
		Type:     notificationReq.Type,
		IsRead:   entity.Unread,
		CreateAt: time.Now(),
	}
	err := ns.notificationRepo.AddNotification(notification)
	if err != nil {
		return err
	}
	return nil
}

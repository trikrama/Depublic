package service

import (
	"context"

	"github.com/trikrama/Depublic/internal/app/notification/entity"
	"github.com/trikrama/Depublic/internal/app/notification/repository"
)

type NotificationServiceInterface interface {
	CreateNotification(c context.Context, notification *entity.Notification) error
	UpdateNotification(c context.Context, notification *entity.Notification) error
	DeleteNotification(c context.Context, id int) error
	GetByUser(c context.Context, id int64) ([]*entity.Notification, error)
	GetAllNotifications(c context.Context) ([]*entity.Notification, error)
}

type NotificationService struct {
	repo repository.NotificationRepositoryInterface
}

func NewNotificationService(repo repository.NotificationRepositoryInterface) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) CreateNotification(c context.Context, notification *entity.Notification) error {
	return s.repo.CreateNotification(c, notification)
}

func (s *NotificationService) UpdateNotification(c context.Context, notification *entity.Notification) error {
	return s.repo.UpdateNotification(c, notification)
}

func (s *NotificationService) DeleteNotification(c context.Context, id int) error {
	return s.repo.DeleteNotification(c, id)
}

func (s *NotificationService) GetByUser(c context.Context, id int64) ([]*entity.Notification, error) {
	s.repo.UpdateStatusNotification(c, id, true)
	return s.repo.GetByUser(c, id)
}

func (s *NotificationService) GetAllNotifications(c context.Context) ([]*entity.Notification, error) {
	return s.repo.GetAllNotifications(c)
}

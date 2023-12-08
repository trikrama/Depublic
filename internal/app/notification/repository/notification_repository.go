package repository

import (
	"context"

	"github.com/trikrama/Depublic/internal/app/notification/entity"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	GetByUser(c context.Context, id int64) ([]*entity.Notification, error)
	CreateNotification(c context.Context, notification *entity.Notification) error
	UpdateNotification(c context.Context, notification *entity.Notification) error
	DeleteNotification(c context.Context, id int) error
	GetAllNotifications(c context.Context) ([]*entity.Notification, error)
}

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetByUser(c context.Context, id int64) ([]*entity.Notification, error) {
	notifications := make([]*entity.Notification, 0)
	err := r.db.WithContext(c).Where("user_id = ?", id).Order("created_at DESC").Limit(15).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) CreateNotification(c context.Context, notification *entity.Notification) error {
	err := r.db.WithContext(c).Create(&notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepository) UpdateNotification(c context.Context, notification *entity.Notification) error {
	err := r.db.WithContext(c).Model(&entity.Notification{}).Where("id = ?", notification.ID).Updates(&notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepository) DeleteNotification(c context.Context, id int) error {
	err := r.db.WithContext(c).Delete(&entity.Notification{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepository) GetAllNotifications(c context.Context) ([]*entity.Notification, error) {
	notifications := make([]*entity.Notification, 0)
	err := r.db.WithContext(c).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
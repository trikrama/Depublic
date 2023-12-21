package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/trikrama/Depublic/internal/app/notification/entity"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	GetByUser(c context.Context, id int64) ([]*entity.Notification, error)
	CreateNotification(c context.Context, notification *entity.Notification) error
	UpdateNotification(c context.Context, notification *entity.Notification) error
	DeleteNotification(c context.Context, id int) error
	GetAllNotifications(c context.Context) ([]*entity.Notification, error)
	UpdateStatusNotification(c context.Context, id int64, status bool) error
}

const (
	NotificationKey = "notifications:all"
)

type NotificationRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewNotificationRepository(db *gorm.DB, redisClient *redis.Client) *NotificationRepository {
	return &NotificationRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *NotificationRepository) GetAllNotifications(c context.Context) ([]*entity.Notification, error) {
	notifications := make([]*entity.Notification, 0)
	val, err := r.redisClient.Get(context.Background(), NotificationKey).Result()
	if err != nil {
		err := r.db.WithContext(c).Find(&notifications).Error // SELECT * FROM users
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(notifications)
		if err != nil {
			return nil, err
		}
		err = r.redisClient.Set(c, NotificationKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return notifications, nil
	}
	err = json.Unmarshal([]byte(val), &notifications)
	if err != nil {
		return nil, err
	}
	return notifications, nil
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

func (r *NotificationRepository) UpdateStatusNotification(c context.Context, id int64, status bool) error {
	err := r.db.WithContext(c).Model(&entity.Notification{}).Where("user_id = ?", id).Update("is_read", status).Error
	if err != nil {
		return err
	}
	return nil
}

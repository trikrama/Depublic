package entity

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	Title     string         `json:"title"`
	Body      string         `json:"content"`
	Status    string         `json:"status"`
	IsRead    bool           `json:"is_read"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func NewNotification(n NotificationRequest) *Notification {
	return &Notification{
		UserID: n.UserID,
		Title:  n.Title,
		Body:   n.Body,
		Status: n.Status,
	}
}

func NewNotificationUpdate(n NotificationRequestUpdate) *Notification {
	return &Notification{
		ID:    n.ID,
		Title: n.Title,
		Body:  n.Message,
	}
}

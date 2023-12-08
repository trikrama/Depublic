package entity

import "time"

type NotificationResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNotificationResponse(n *Notification) *NotificationResponse {
	return &NotificationResponse{
		ID:        n.ID,
		Title:     n.Title,
		Content:   n.Body,
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	}
}

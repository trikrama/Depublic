package entity

type NotificationRequest struct {
	UserID  int64  `json:"user_id" validate:"required"`
	Title   string `json:"title"`
	Body string `json:"message"`
	Status  string `json:"status"`
}

type NotificationRequestUpdate struct {
	ID      int64  `param:"id" validate:"required"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
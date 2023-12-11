package entity

import "github.com/google/uuid"

type TransactionResponse struct {
	ID        uuid.UUID `json:"id"`
	EventID   int64     `json:"event_id"`
	UserID    int64     `json:"user_id"`
	Quantity  int       `json:"quantity"`
	Total     int       `json:"total"`
	Redirect  string    `json:"redirect_url"`
	Status    string    `json:"status"`
	CreatedAt string    `json:"created_at"`
}

func NewTransactionResponse(t *Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:        t.ID,
		EventID:   t.EventID,
		UserID:    t.UserID,
		Quantity:  t.Quantity,
		Total:     t.Total,
		Redirect:  t.RedirectURL,
		Status:    t.TransactionStatus,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

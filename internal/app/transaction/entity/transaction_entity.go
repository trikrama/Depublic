package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                uuid.UUID    `json:"id"`
	UserID            int64      `json:"userID"`
	EventID           int64      `json:"eventID"`
	TransactionStatus string     `json:"transaction_status"` // pending, success, failed
	Quantity          int        `json:"quantity"`
	Total             int        `json:"total"`
	RedirectURL       string     `json:"redirect_url" gorm:"-"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	DeletedAt         *time.Time `json:"deletedAt"`
}

func NewTransaction(t TransactionRequest) *Transaction {
	return &Transaction{
		UserID:            t.UserId,
		EventID:           t.EventId,
		Quantity:          t.Quantity,
		CreatedAt:         time.Now(),
	}
}

func NewTransactionUpdate(t TransactionRequestUpdate) *Transaction {
	return &Transaction{
		ID:                t.ID,
		UserID:            t.UserId,
		EventID:           t.EventId,
		Quantity:          t.Quantity,
		UpdatedAt:         time.Now(),
	}
}

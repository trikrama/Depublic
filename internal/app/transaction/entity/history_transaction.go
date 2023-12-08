package entity

import (
	"time"

	"github.com/google/uuid"
)

type HistoryTransaction struct {
	HistoryId     int64     `json:"history_id" gorm:"primaryKey;autoIncrement"`
	TransactionID uuid.UUID `json:"transaction_id"`
	UserID        int64     `json:"user_id"`
	Action        string    `json:"action"`
	Timestamp     time.Time `json:"timestamp"`
	NameEvent     string    `json:"name_event"`
	Quantity      int64     `json:"quantity"`
	Total         int64     `json:"total"`
}

func NewHistoryTransaction(
	transactionID uuid.UUID,
	userID int64,
	action string,
	timestamp time.Time,
	nameEvent string,
	quantity int64,
	total int64) HistoryTransaction {
	return HistoryTransaction{
		TransactionID: transactionID,
		UserID:        userID,
		Action:        action,
		Timestamp:     timestamp,
		NameEvent:     nameEvent,
		Quantity:      quantity,
		Total:         total,
	}
}

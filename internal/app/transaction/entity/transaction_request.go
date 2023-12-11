package entity

import "github.com/google/uuid"

type TransactionRequest struct {
	UserId   int64 `json:"userID"`
	EventId  int64 `json:"eventID" validate:"required"`
	Quantity int   `json:"quantity" validate:"required"`
}

type TransactionRequestUpdate struct {
	ID       uuid.UUID `param:"id" validate:"required"`
	UserId   int64 `json:"userID" validate:"required"`
	EventId  int64 `json:"eventID" validate:"required"`
	Quantity int   `json:"quantity" validate:"required"`
}
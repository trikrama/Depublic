package entity

import "github.com/google/uuid"

type TransactionPaymentMidtrans struct {
	VaNumbers []struct {
		VaNumber string `json:"va_number"`
		Bank     string `json:"bank"`
	} `json:"va_numbers"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	TransactionID     string        `json:"transaction_id"`
	StatusMessage     string        `json:"status_message"`
	StatusCode        string        `json:"status_code"`
	SignatureKey      string        `json:"signature_key"`
	SettlementTime    string        `json:"settlement_time"`
	PaymentType       string        `json:"payment_type"`
	PaymentAmounts    []interface{} `json:"payment_amounts"`
	OrderID           string        `json:"order_id"`
	MerchantID        string        `json:"merchant_id"`
	GrossAmount       string        `json:"gross_amount"`
	FraudStatus       string        `json:"fraud_status"`
	ExpiryTime        string        `json:"expiry_time"`
	Currency          string        `json:"currency"`
}

type PaymentRequest struct {
	OrderID   uuid.UUID
	Amount    int64
	Name      string
	Email     string
}

func NewPaymentRequest(orderID uuid.UUID, amount int64, Name, email string) *PaymentRequest {
	return &PaymentRequest{
		OrderID:   orderID,
		Amount:    amount,
		Name:      Name,
		Email:     email,
	}
}

// type PaymentRequest struct {
// 	OrderID uuid.UUID
// 	EventID int64
// 	EventName string
// 	PriceEvent int64
// 	EventQuantity int64
// 	TransactionQuantity int
// 	Amount  int64
// 	Status  string
// 	Name    string
// 	Email   string
// }

// func NewPaymentRequest(orderId uuid.UUID, eventId, priceEvent, Amount, eventQuantity int64, transactionQuantity int, eventName, status, name, email string) *PaymentRequest {
// 	return &PaymentRequest{
// 		OrderID: orderId,
// 		EventID: eventId,
// 		EventName: eventName,
// 		PriceEvent: priceEvent,
// 		EventQuantity: eventQuantity,
// 		Status: status,
// 		TransactionQuantity: transactionQuantity,
// 		Amount:  Amount,
// 		Name:  name,
// 		Email: email,
// 	}
// }

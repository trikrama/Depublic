package service

import (
	"context"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
)

type PaymentServiceInterface interface {
	CreateTransactionMidtrans(ctx context.Context, paymentRequest *entity.PaymentRequest) (string, error)
}

type PaymentService struct {
	client snap.Client
}

func NewPaymentService(client snap.Client) *PaymentService {
	return &PaymentService{
		client: client,
	}
}

func (s *PaymentService) CreateTransactionMidtrans(ctx context.Context, paymentRequest *entity.PaymentRequest) (string, error) {
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentRequest.OrderID.String(),
			GrossAmt: paymentRequest.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: paymentRequest.Name,
			Email: paymentRequest.Email,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.Itoa(int(paymentRequest.EventID)),
				Name:  paymentRequest.Name,
				Qty:   int32(paymentRequest.TransactionQuantity),
				Price: paymentRequest.PriceEvent,
			},
		},
	}
	snapResponse, err := s.client.CreateTransaction(request)
	if err != nil {
		return "", err
	}
	return snapResponse.RedirectURL, nil
}

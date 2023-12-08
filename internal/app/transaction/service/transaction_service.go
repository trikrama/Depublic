package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/snap"

	// repositoryEvent "github.com/trikrama/Depublic/internal/app/event/repository"
	entityEvent "github.com/trikrama/Depublic/internal/app/event/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/repository"

	// repositoryUser "github.com/trikrama/Depublic/internal/app/user/repository"
	"github.com/trikrama/Depublic/internal/config"
)

type TransactionServiceInterface interface {
	CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	UpdateTransaction(c context.Context, transaction *entity.Transaction) error
	DeleteTransaction(c context.Context, id int) error
	GetAllTransaction(c context.Context) ([]*entity.Transaction, error)
	GetTransactionByID(c context.Context, id string) (*entity.Transaction, error)
	GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error)
	UpdateStatus(c context.Context, id uuid.UUID, status string) error
	GetEvent(c context.Context, idEvent int64) (*entityEvent.Event, error)
	CreateHistory(c context.Context, history *entity.HistoryTransaction) error
	GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error)
	GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error)
}

type TransactionService struct {
	repo repository.TransactionRepositoryInterface
	// repoEvent repositoryEvent.EventRepositoryInterface
	// repoUser  repositoryUser.UserRepositoryInterface
}

func NewTransactionService(repo repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	event, err := s.repo.GetEvent(c, transaction.EventID)
	if err != nil {
		return nil, errors.New("event not found")
	}

	user, err := s.repo.GetUserById(c, transaction.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if event.Status != "Available" {
		return nil, errors.New("event not availebale")
	}

	if event.Quantity < int64(transaction.Quantity) {
		return nil, errors.New("tiket sudah habis")
	}

	total := event.Price * int64(transaction.Quantity)
	transaction.Total = int(total)

	var cfg *config.Config
	cfg, _ = config.NewConfig(".env")
	serverKey := cfg.Midtrans.ServerKey

	transaction.ID = uuid.New()
	var snapClient = snap.Client{}
	snapClient.New(serverKey, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.ID.String(),
			GrossAmt: int64(transaction.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			Phone: user.Number,
		},
		Items: &[]midtrans.ItemDetails{

			{
				ID:    strconv.Itoa(int(event.ID)),
				Price: int64(event.Price),
				Qty:   int32(transaction.Quantity),
				Name:  event.Name,
			},
		},
	}

	res, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		fmt.Println("error 5")
		return nil, err
	}

	event.Quantity -= int64(transaction.Quantity)
	err = s.repo.UpdateEvent(c, event.ID, event.Quantity)
	if err != nil {
		return nil, err
	}
	if err != nil {
		fmt.Println("error 4")
		return nil, err
	}

	transaction.RedirectURL = res.RedirectURL

	transaction.TransactionStatus = "Pending"

	tx, err1 := s.repo.CreateTransaction(c, transaction)

	if err1 != nil {
		fmt.Println("error 6")
		return nil, err1
	}

	return tx, nil
}

func (s *TransactionService) UpdateTransaction(c context.Context, transaction *entity.Transaction) error {
	event, err := s.repo.GetEvent(c, transaction.EventID)
	if err != nil {
		return errors.New("event not found")
	}
	if event.Status != "Available" {
		return errors.New("event not availebale")
	}
	if event.Quantity < int64(transaction.Quantity) {
		return errors.New("tiket sudah habis")
	}
	total := event.Price * int64(transaction.Quantity)
	transaction.Total = int(total)
	return s.repo.UpdateTransaction(c, transaction)
}

func (s *TransactionService) DeleteTransaction(c context.Context, id int) error {
	return s.repo.DeleteTransaction(c, id)
}

func (s *TransactionService) GetAllTransaction(c context.Context) ([]*entity.Transaction, error) {
	return s.repo.GetAllTransaction(c)
}

func (s *TransactionService) GetTransactionByID(c context.Context, id string) (*entity.Transaction, error) {
	return s.repo.GetTransactionByID(c, id)
}

func (s *TransactionService) GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error) {
	return s.repo.GetTransactionByUser(c, id)
}

func (s *TransactionService) UpdateStatus(c context.Context, id uuid.UUID, status string) error {
	return s.repo.UpdateStatusTransaction(c, id, status)
}

func (s *TransactionService) GetEvent(c context.Context, idEvent int64) (*entityEvent.Event, error) {
	return s.repo.GetEvent(c, idEvent)
}

func (s *TransactionService) CreateHistory(c context.Context, history *entity.HistoryTransaction) error {
	return s.repo.CreateHistory(c, history)
}

func (s *TransactionService) GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error) {
	return s.repo.GetAllHistory(c)
}

func (s *TransactionService) GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error) {
	return s.repo.GetHistoryByUser(c, id)
}

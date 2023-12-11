package service

import (
	"context"
	"errors"
	"fmt"
 
	// repositoryEvent "github.com/trikrama/Depublic/internal/app/event/repository"
	entityEvent "github.com/trikrama/Depublic/internal/app/event/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/repository"
	// repositoryUser "github.com/trikrama/Depublic/internal/app/user/repository"
)

type TransactionServiceInterface interface {
	CreateTransaction(c context.Context, transaction *entity.Transaction, event *entityEvent.Event) (*entity.Transaction, error)
	UpdateTransaction(c context.Context, transaction *entity.Transaction, event *entityEvent.Event) error
	DeleteTransaction(c context.Context, id int) error
	GetAllTransaction(c context.Context) ([]*entity.Transaction, error)
	GetTransactionByID(c context.Context, id string) (*entity.Transaction, error)
	GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error)
	CreateHistory(c context.Context, history *entity.HistoryTransaction) error
	GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error)
	GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error)
}

type TransactionService struct {
	repo repository.TransactionRepositoryInterface
}

func NewTransactionService(repo repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(c context.Context, transaction *entity.Transaction, event *entityEvent.Event) (*entity.Transaction, error) {
	if event.Status != "Available" {
		return nil, errors.New("event not availebale")
	}

	if event.Quantity < int64(transaction.Quantity) {
		return nil, fmt.Errorf("There are only %d tickets available", event.Quantity)
	}



	total := event.Price * int64(transaction.Quantity)
	transaction.Total = int(total)

	tx, err1 := s.repo.CreateTransaction(c, transaction)

	if err1 != nil {
		fmt.Println("error 6")
		return nil, err1
	}

	return tx, nil
}

func (s *TransactionService) UpdateTransaction(c context.Context, transaction *entity.Transaction, event *entityEvent.Event) error {
	tx, err := s.GetTransactionByID(c, transaction.ID.String())
	if err != nil {
		return errors.New("transaction not found")
	}
	if event.Status != "Available" {
		return errors.New("event not availebale")
	}
	if event.Quantity < int64(transaction.Quantity) {
		return fmt.Errorf("There are only %d tickets available", event.Quantity)
	}
	transaction.TransactionStatus = tx.TransactionStatus
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

func (s *TransactionService) CreateHistory(c context.Context, history *entity.HistoryTransaction) error {
	return s.repo.CreateHistory(c, history)
}

func (s *TransactionService) GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error) {
	return s.repo.GetAllHistory(c)
}

func (s *TransactionService) GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error) {
	return s.repo.GetHistoryByUser(c, id)
}

 
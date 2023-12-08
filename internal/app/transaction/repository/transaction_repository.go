package repository

import (
	"context"

	"github.com/google/uuid"
	entityEvent "github.com/trikrama/Depublic/internal/app/event/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
	entityUser "github.com/trikrama/Depublic/internal/app/user/entity"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	GetAllTransaction(c context.Context) ([]*entity.Transaction, error)
	GetTransactionByID(c context.Context, id string) (*entity.Transaction, error)
	CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	UpdateTransaction(c context.Context, transaction *entity.Transaction) error
	DeleteTransaction(c context.Context, id int) error
	GetEvent(c context.Context, idEvent int64) (*entityEvent.Event, error)
	GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error)
	GetUserById(c context.Context, id int64) (*entityUser.User, error)
	UpdateStatusTransaction(c context.Context, id uuid.UUID, status string) error
	UpdateEvent(c context.Context, id int64, quantity int64) error
	CreateHistory(c context.Context, history *entity.HistoryTransaction) error 
	GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error)
	GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}


func (t *TransactionRepository) GetAllTransaction(c context.Context) ([]*entity.Transaction, error) {
	transaction := make([]*entity.Transaction, 0)
	err := t.db.WithContext(c).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}


func (t *TransactionRepository) GetTransactionByID(c context.Context, id string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	err := t.db.WithContext(c).First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}


func (t *TransactionRepository) GetEvent(c context.Context, idEvent int64) (*entityEvent.Event, error) {
	event := new(entityEvent.Event)
	err := t.db.WithContext(c).First(&event, idEvent).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (t *TransactionRepository) CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	err := t.db.WithContext(c).Create(&transaction).Error
	if err != nil {
		return &entity.Transaction{}, err
	}
	return transaction, nil
}



func (t *TransactionRepository) UpdateTransaction(c context.Context, transaction *entity.Transaction) error {
	err := t.db.WithContext(c).Model(&entity.Transaction{}).Where("id = ?", transaction.ID).Updates(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}


func (t *TransactionRepository) DeleteTransaction(c context.Context, id int) error {
	err := t.db.WithContext(c).Delete(&entity.Transaction{}, id).Error
	if err != nil {
		return err
	}
	return nil
}



func (t *TransactionRepository) GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error) {
	transaction := make([]*entity.Transaction, 0)
	err := t.db.WithContext(c).Where("user_id = ?", id).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}



func (t *TransactionRepository) GetUserById(c context.Context, id int64) (*entityUser.User, error) {
	user := new(entityUser.User)
	err := t.db.WithContext(c).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (t *TransactionRepository) UpdateStatusTransaction(c context.Context, id uuid.UUID, status string) error {
	// idTransaction, _ := uuid.Parse(id)
	err := t.db.WithContext(c).Model(&entity.Transaction{}).Where("id = ?", id).Update("transaction_status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) UpdateEvent(c context.Context, id int64, quantity int64) error{
	err := t.db.WithContext(c).Model(&entityEvent.Event{}).Where("id = ?", id).Update("quantity", quantity).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) CreateHistory(c context.Context, history *entity.HistoryTransaction) error {
	err := t.db.WithContext(c).Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error) {
	history := make([]*entity.HistoryTransaction, 0)
	err := t.db.WithContext(c).Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}


func (t *TransactionRepository) GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error) {
	history := make([]*entity.HistoryTransaction, 0)
	err := t.db.WithContext(c).Where("user_id = ?", id).Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}
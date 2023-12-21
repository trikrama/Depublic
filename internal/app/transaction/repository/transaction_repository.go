package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	GetAllTransaction(c context.Context) ([]*entity.Transaction, error)
	GetTransactionByID(c context.Context, id string) (*entity.Transaction, error)
	CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	UpdateTransaction(c context.Context, transaction *entity.Transaction) error
	DeleteTransaction(c context.Context, id uuid.UUID) error
	GetTransactionByUser(c context.Context, id int64) ([]*entity.Transaction, error)
	CreateHistory(c context.Context, history *entity.HistoryTransaction) error
	GetAllHistory(c context.Context) ([]*entity.HistoryTransaction, error)
	GetHistoryByUser(c context.Context, id int64) ([]*entity.HistoryTransaction, error)
}

const (
	TransactionKey = "transactions:all"
)

type TransactionRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewTransactionRepository(db *gorm.DB, redisClient *redis.Client) *TransactionRepository {
	return &TransactionRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (t *TransactionRepository) GetAllTransaction(c context.Context) ([]*entity.Transaction, error) {
	transaction := make([]*entity.Transaction, 0)
	val, err := t.redisClient.Get(context.Background(), TransactionKey).Result()
	if err != nil {
		err := t.db.WithContext(c).Find(&transaction).Error
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(transaction)
		if err != nil {
			return nil, err
		}
		err = t.redisClient.Set(c, TransactionKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return transaction, nil	
	}
	err = json.Unmarshal([]byte(val), &transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionRepository) GetTransactionByID(c context.Context, id string) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	err := t.db.WithContext(c).Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionRepository) CreateTransaction(c context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	transaction.ID = uuid.New()
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

func (t *TransactionRepository) DeleteTransaction(c context.Context, id uuid.UUID) error {
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

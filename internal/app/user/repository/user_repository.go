package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/trikrama/Depublic/internal/app/user/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAllUser(c context.Context) ([]*entity.User, error)
	GetUserByID(c context.Context, id int) (*entity.User, error)
	CreateUser(c context.Context, user *entity.User) error
	UpdateUser(c context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(c context.Context, id int) error
	GetUserByEmail(c context.Context, email string) (*entity.User, error)
}

const (
	UserKey = "users:all"
)

type UserRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewRepositoryUser(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *UserRepository) GetAllUser(c context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	val, err := r.redisClient.Get(context.Background(), UserKey).Result()
	if err != nil {
		err := r.db.WithContext(c).Find(&users).Error // SELECT * FROM users
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(users)
		if err != nil {
			return nil, err
		}

		// Set the data in Redis with an expiration time (e.g., 1 hour)
		err = r.redisClient.Set(c, UserKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(c context.Context, id int) (*entity.User, error) {
	user := new(entity.User)
	tx := r.db.WithContext(c).First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) CreateUser(c context.Context, user *entity.User) error {
	err := r.db.WithContext(c).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(c context.Context, user *entity.User) (*entity.User, error) {
	err := r.db.WithContext(c).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(c context.Context, id int) error {
	err := r.db.WithContext(c).Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(c context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

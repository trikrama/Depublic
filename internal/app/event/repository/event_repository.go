package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/trikrama/Depublic/internal/app/event/entity"
	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	GetAllEvent(c context.Context, queryFilter entity.QueryFilter) ([]*entity.Event, error)
	GetEventByID(c context.Context, id int64) (*entity.Event, error)
	CreateEvent(c context.Context, event *entity.Event) error
	UpdateEvent(c context.Context, event *entity.Event) error
	DeleteEvent(c context.Context, id int) error
}

const (
	EventKey = "events:all"
)

type EventRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewEventRepository(db *gorm.DB, redisClient *redis.Client) *EventRepository {
	return &EventRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *EventRepository) GetAllEvent(c context.Context, queryFilter entity.QueryFilter) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	val, err := r.redisClient.Get(context.Background(), EventKey).Result()
	if err != nil {
		db := r.db.WithContext(c)

		// Filter berdasarkan harga
		if queryFilter.Filter.Price > 0 {
			db = db.Where("price = ?", queryFilter.Filter.Price)
		}

		// Filter berdasarkan lokasi
		if queryFilter.Filter.Location != "" {
			db = db.Where("location = ?", queryFilter.Filter.Location)
		}

		// Filter berdasarkan tanggal
		if queryFilter.Filter.StartDate != "" && queryFilter.Filter.EndDate != "" {
			startDate, err := time.Parse("2006-01-02", queryFilter.Filter.StartDate)
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse("2006-01-02", queryFilter.Filter.EndDate)
			if err != nil {
				return nil, err
			}
			db = db.Where("date BETWEEN ? AND ?", startDate, endDate)
		}

		// Filter berdasarkan keyword
		if queryFilter.Search != "" {
			keyword := fmt.Sprintf("%%%s%%", queryFilter.Search)
			db = db.Where("name LIKE ? OR location LIKE ?", keyword, keyword)
		}

		// Filter berdasarkan status
		if queryFilter.Filter.Status != "" {
			db = db.Where("status = ?", queryFilter.Filter.Status)
		}

		// Sorting
		if queryFilter.Sort.By != "" && queryFilter.Sort.Order != "" {
			order := fmt.Sprintf("%s %s", queryFilter.Sort.By, queryFilter.Sort.Order)
			db = db.Order(order)
		}
		err := db.Find(&events).Error
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(events)
		if err != nil {
			return nil, err
		}
		// Set the data in Redis with an expiration time (e.g., 1 hour)
		err = r.redisClient.Set(c, EventKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return events, nil
	}
	err = json.Unmarshal([]byte(val), &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetEventByID(c context.Context, id int64) (*entity.Event, error) {
	event := new(entity.Event)
	err := r.db.WithContext(c).Where("id = ?", id).First(&event).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) CreateEvent(c context.Context, event *entity.Event) error {
	err := r.db.WithContext(c).Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) UpdateEvent(c context.Context, event *entity.Event) error {
	err := r.db.WithContext(c).Model(&entity.Event{}).Where("id = ?", event.ID).Updates(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) DeleteEvent(c context.Context, id int) error {
	err := r.db.WithContext(c).Delete(&entity.Event{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

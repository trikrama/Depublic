package repository

import (
	"context"
	"errors"
	"time"

	"github.com/trikrama/Depublic/internal/app/event/entity"
	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	GetAllEvent(c context.Context) ([]*entity.Event, error)
	GetEventByID(c context.Context, id int) (*entity.Event, error)
	CreateEvent(c context.Context, event *entity.Event) error
	UpdateEvent(c context.Context, event *entity.Event) error
	DeleteEvent(c context.Context, id int) error
	SearchEvent(ctx context.Context, search string) ([]*entity.Event, error)
	SortEventByNewest(ctx context.Context) ([]*entity.Event, error)
	SortEventByCheapest(ctx context.Context) ([]*entity.Event, error)
	SortEventByExpensive(ctx context.Context) ([]*entity.Event, error)
	SortEventByStatus(ctx context.Context, status string) ([]*entity.Event, error)
	FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error)
	FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error)
	FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error)
}

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetAllEvent(c context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(c).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetEventByID(c context.Context, id int) (*entity.Event, error) {
	event := new(entity.Event)
	err := r.db.WithContext(c).First(&event, id).Error
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
	tx := r.db.WithContext(c).Where("id = ?", event.ID).First(&event)
	if tx.Error != nil {
		return tx.Error
	}
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

// SearchEvent
func (r *EventRepository) SearchEvent(ctx context.Context, search string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	result := r.db.WithContext(ctx).Where("name LIKE ?", "%"+search+"%").Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func (r *EventRepository) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("price >= ? And price <= ?", min, max).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("start_date >= ? And end_date <= ?", startDate, endDate).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("location ILIKE ?", "%"+location+"%").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByStatus(ctx context.Context, status string) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Where("status = ?", status).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByExpensive(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("price DESC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByCheapest(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("price ASC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) SortEventByNewest(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/trikrama/Depublic/internal/app/event/entity"
	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	GetAllEvent(c context.Context) ([]*entity.Event, error)
	GetEventByID(c context.Context, id int64) (*entity.Event, error)
	CreateEvent(c context.Context, event *entity.Event) error
	UpdateEvent(c context.Context, event *entity.Event) error
	DeleteEvent(c context.Context, id int) error
	GetByFilter(c context.Context, queryFilter entity.QueryFilter) ([]*entity.Event, error)
	// SearchEvent(ctx context.Context, search string) ([]*entity.Event, error)
	// SortEventByNewest(ctx context.Context) ([]*entity.Event, error)
	// SortEventByCheapest(ctx context.Context) ([]*entity.Event, error)
	// SortEventByExpensive(ctx context.Context) ([]*entity.Event, error)
	// SortEventByStatus(ctx context.Context, status string) ([]*entity.Event, error)
	// FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error)
	// FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error)
	// FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error)
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

// FindByFilter mengembalikan event yang difilter berdasarkan query filter
func (r *EventRepository) GetByFilter(c context.Context, queryFilter entity.QueryFilter) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0)
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
	return events, nil
}

// // SearchEvent
// func (r *EventRepository) SearchEvent(ctx context.Context, search string) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	result := r.db.WithContext(ctx).Where("name LIKE ?", "%"+search+"%").Find(&events)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return events, nil
// }

// func (r *EventRepository) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Where("price >= ? And price <= ?", min, max).Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Where("start_date >= ? And end_date <= ?", startDate, endDate).Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Where("location ILIKE ?", "%"+location+"%").Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) SortEventByStatus(ctx context.Context, status string) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Where("status = ?", status).Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) SortEventByExpensive(ctx context.Context) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Order("price DESC").Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) SortEventByCheapest(ctx context.Context) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Order("price ASC").Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (r *EventRepository) SortEventByNewest(ctx context.Context) ([]*entity.Event, error) {
// 	events := make([]*entity.Event, 0)
// 	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&events).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

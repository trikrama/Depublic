package service

import (
	"context"

	"github.com/trikrama/Depublic/internal/app/event/entity"
	"github.com/trikrama/Depublic/internal/app/event/repository"
)

type EventServiceInterface interface {
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

type EventService struct {
	repo repository.EventRepositoryInterface
}

func NewEventService(repo repository.EventRepositoryInterface) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) GetAllEvent(c context.Context) ([]*entity.Event, error) {
	return s.repo.GetAllEvent(c)
}

func (s *EventService) GetEventByID(c context.Context, id int64) (*entity.Event, error) {
	return s.repo.GetEventByID(c, id)
}

func (s *EventService) CreateEvent(c context.Context, event *entity.Event) error {
	return s.repo.CreateEvent(c, event)
}

func (s *EventService) UpdateEvent(c context.Context, event *entity.Event) error {
	return s.repo.UpdateEvent(c, event)
}

func (s *EventService) DeleteEvent(c context.Context, id int) error {
	return s.repo.DeleteEvent(c, id)
}

func (s *EventService) GetByFilter(c context.Context, queryFilter entity.QueryFilter) ([]*entity.Event, error) {
	return s.repo.GetByFilter(c, queryFilter)
}

// func (s *EventService) SearchEvent(ctx context.Context, search string) ([]*entity.Event, error) {
// 	return s.repo.SearchEvent(ctx, search)
// }

// func (s *EventService) SortEventByNewest(ctx context.Context) ([]*entity.Event, error) {
// 	return s.repo.SortEventByNewest(ctx)
// }

// func (s *EventService) SortEventByCheapest(ctx context.Context) ([]*entity.Event, error) {
// 	return s.repo.SortEventByCheapest(ctx)
// }

// func (s *EventService) SortEventByExpensive(ctx context.Context) ([]*entity.Event, error) {
// 	return s.repo.SortEventByExpensive(ctx)
// }

// func (s *EventService) SortEventByStatus(ctx context.Context, status string) ([]*entity.Event, error) {
// 	return s.repo.SortEventByStatus(ctx, status)
// }

// func (s *EventService) FilterEventByPrice(ctx context.Context, min, max string) ([]*entity.Event, error) {
// 	return s.repo.FilterEventByPrice(ctx, min, max)
// }

// func (s *EventService) FilterEventByLocation(ctx context.Context, location string) ([]*entity.Event, error) {
// 	return s.repo.FilterEventByLocation(ctx, location)
// }

// func (s *EventService) FilterEventByDate(ctx context.Context, startDate, endDate time.Time) ([]*entity.Event, error) {
// 	return s.repo.FilterEventByDate(ctx, startDate, endDate)
// }

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Image       string         `json:"image"`
	Description string         `json:"description"`
	Category    string         `json:"category"` //music, sport, etc
	Location    string         `json:"lokasi"`
	Price       int64          `json:"price"`
	DateStart   time.Time      `json:"data_start"`
	DateEnd     time.Time      `json:"data_end"`
	Quantity    int64          `json:"quantity"`
	Status      string         `json:"status"` //Tersedia, Tidak Tersedia
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func NewEvent(e EventRequest) *Event {
	return &Event{
		Name:        e.Name,
		Image:       e.Image,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		DateStart:   e.DateStart,
		DateEnd:     e.DateEnd,
		Price:       e.Price,
		Quantity:    e.Quantity,
		Status:      e.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewEventUpdate(e EventRequestUpdate) *Event {
	return &Event{
		ID:          e.ID,
		Name:        e.Name,
		Image:       e.Image,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		DateStart:   e.DateStart,
		DateEnd:     e.DateEnd,
		Price:       e.Price,
		Quantity:    e.Quantity,
		Status:      e.Status,
		UpdatedAt:   time.Now(),
	}
}

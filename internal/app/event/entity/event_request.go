package entity

import (
	"time"
)

type EventRequest struct {
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Description string `json:"description"`
	Category  string    `json:"category"`
	Location  string    `json:"lokasi"`
	DateStart time.Time `json:"data_start"`
	DateEnd   time.Time `json:"data_end"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity"`
	Status    string    `json:"status"`
}

type EventRequestUpdate struct {
	ID       int64     `param:"id" validate:"required"`
	Name     string    `json:"name"`
	Image    string    `json:"image"`
	Description string `json:"description"`
	Category string    `json:"category"`
	Location string    `json:"lokasi"`
	DateStart time.Time `json:"data_start"`
	DateEnd   time.Time `json:"data_end"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity"`
	Status   string    `json:"status"`
}

// func NewEventRequest(e Event) *Event {
// 	return &Event{
// 		Name:     e.Name,
// 		Category: e.Category,
// 		Location: e.Location,
// 		Price:    e.Price,
// 		Date:     e.Date,
// 		Status:   e.Status,
// 	}
// }

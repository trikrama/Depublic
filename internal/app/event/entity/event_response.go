package entity

type EventResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Location    string `json:"lokasi"`
	DateStart   string `json:"data_start"`
	DateEnd     string `json:"data_end"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Status      string `json:"status"`
}

func NewEventResponse(e *Event) *EventResponse {
	return &EventResponse{
		ID:          e.ID,
		Name:        e.Name,
		Image:       e.Image,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		DateStart:   e.DateStart.Format("02 Januari 2006"),
		DateEnd:     e.DateEnd.Format("02 Januari 2006"),
		Price:       e.Price,
		Quantity:    e.Quantity,
		Status:      e.Status,
	}
}

package entity

type EventResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Description string `json:"description"`
	Category  string `json:"category"`
	Location  string `json:"lokasi"`
	DateStart string `json:"data_start"`
	DateEnd   string `json:"data_end"`
	Price     int64  `json:"price"`
	Quantity  int64  `json:"quantity"`
	Status    string `json:"status"`
}

func NewEventResponse(e *Event) *EventResponse {
	return &EventResponse{
		ID:        e.ID,
		Name:      e.Name,
		Image:     e.Image,
		Description: e.Description,
		Category:  e.Category,
		Location:  e.Location,
		DateStart: e.DateStart.Format("2006-01-02"),
		DateEnd:   e.DateEnd.Format("2006-01-02"),
		Price:     e.Price,
		Quantity:  e.Quantity,
		Status:    e.Status,
	}
}

type EventBeforeLogin struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Location  string `json:"lokasi"`
	DateStart string `json:"data_start"`
	DateEnd   string `json:"data_end"`
	Price     int64  `json:"price"`
	Status    string `json:"status"`
}

func NewEventBeforeLogin(e *Event) *EventBeforeLogin {
	return &EventBeforeLogin{
		Name:      e.Name,
		Category:  e.Category,
		Location:  e.Location,
		DateStart: e.DateStart.Format("2006-01-02"),
		DateEnd:   e.DateEnd.Format("2006-01-02"),
		Price:     e.Price,
		Status:    e.Status,
	}
}
package entity

type BlogRequest struct {
	Image string `json:"image"`
	Date  string `json:"date"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogRequestUpdate struct {
	ID    int64  `param:"id" validate:"required"`
	Image string `json:"image"`
	Date  string `json:"date"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

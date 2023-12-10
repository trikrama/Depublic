package entity

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        int64          `json:"id"`
	Image     string         `json:"image"`
	Date      string         `json:"date"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func NewBlog(b BlogRequest) *Blog {
	return &Blog{
		Image: b.Image,
		Date:  b.Date,
		Title: b.Title,
		Body:  b.Body,
	}
}

func NewBlogUpdate(b BlogRequestUpdate) *Blog {
	return &Blog{
		ID:    b.ID,
		Image: b.Image,
		Date:  b.Date,
		Title: b.Title,
		Body:  b.Body,
	}
}

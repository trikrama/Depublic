package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Number    string         `json:"number"`
	Password  string         `json:"password"`
	Role      string         `json:"role" default:"buyer"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func NewUser(u UserRequest) *User {
	return &User{
		Name:      u.Name,
		Email:     u.Email,
		Number:    u.Number,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewUserUpdate(u UserRequestUpdate) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Number:    u.Number,
		Password:  u.Password,
		UpdatedAt: time.Now(),
	}
}

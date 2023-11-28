package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func NewUser(u  UserRequest) *User {
	return &User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateUser(id int64, name, email, password, role string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}






package entity


type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Number   string `json:"number"`
	Role     string `json:"role"`
}

func NewUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Number:   user.Number,
		Role:     user.Role,
	}
}
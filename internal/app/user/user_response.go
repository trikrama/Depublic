package user


type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func NewUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}
}
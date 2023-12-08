package entity

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Number   string `json:"number"`
	Password string `json:"password"`
}

type UserRequestUpdate struct {
	ID       int64  `param:"id" validate:"required"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Number   string `json:"number"`
	Password string `json:"password"`
}
 

type UserRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
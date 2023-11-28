package user

import "context"

type UserServiceInterface interface {
	GetAllUser(c context.Context)([]*User, error)
	GetUserByID(c context.Context, id int)(*User, error)
	CreateUser(c context.Context, user *User)error
	UpdateUser(c context.Context, user *User) error
	Delete(c context.Context, id int)error
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService)GetAllUser(c context.Context)([]*User, error){
	return s.repo.GetAllUser(c)
}

func (s *UserService)GetUserByID(c context.Context, id int)(*User, error){
	return s.repo.GetUserByID(c, id)
}

func (s *UserService)CreateUser(c context.Context, user *User)error{
	return s.repo.CreateUser(c, user)
}

func (s *UserService)UpdateUser(c context.Context, user *User) error{
	return s.repo.UpdateUser(c, user)
}

func (s *UserService)Delete(c context.Context, id int)error{
	return s.repo.DeleteUser(c,id)
}
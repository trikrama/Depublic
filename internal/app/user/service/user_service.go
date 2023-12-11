package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/trikrama/Depublic/common"
	"github.com/trikrama/Depublic/helper"
	"github.com/trikrama/Depublic/internal/app/user/entity"
	"github.com/trikrama/Depublic/internal/app/user/repository"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type UserServiceInterface interface {
	GetAllUser(c context.Context) ([]*entity.User, error)
	GetUserByID(c context.Context, id int) (*entity.User, error)
	CreateUser(c context.Context, user *entity.User) error
	UpdateUser(c context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(c context.Context, id int) error
	LoginUser(c context.Context, email string, password string) (*entity.User, string, error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUser(c context.Context) ([]*entity.User, error) {
	return s.repo.GetAllUser(c)
}

func (s *UserService) GetUserByID(c context.Context, id int) (*entity.User, error) {
	return s.repo.GetUserByID(c, id)
}

// Membuat akun
func (s *UserService) CreateUser(c context.Context, user *entity.User) error {
	//validasi data kosong
	errEmpty := validator.CheckDataEmpty(user.Name, user.Email, user.Password)
	if errEmpty != nil {
		return errEmpty
	}
	//validasi email
	errEmail := validator.EmailFormat(user.Email)
	if errEmail != nil {
		return errEmail
	}

	email, _ := s.repo.GetUserByEmail(c, user.Email)
	if email != nil {
		return errors.New("email already used")
	}

	//validasi password
	errLengthPass := validator.MinLength(user.Password, 8)
	if errLengthPass != nil {
		return errLengthPass
	}

	if user.Role == "" {
		user.Role = "Buyer"
	}
	//hash password
	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return errors.New("masalah hashing password")
	}
	user.Password = hashPass

	//simpan user
	return s.repo.CreateUser(c, user)
}

func (s *UserService) UpdateUser(c context.Context, user *entity.User) (*entity.User, error) {
	//validasi data kosong
	errEmpty := validator.CheckDataEmpty(user.Name, user.Email, user.Password)
	if errEmpty != nil {
		return nil, errEmpty
	}
	//validasi email
	errEmail := validator.EmailFormat(user.Email)
	if errEmail != nil {
		return nil, errEmail
	}
	//validasi password
	errLengthPass := validator.MinLength(user.Password, 8)
	if errLengthPass != nil {
		return nil, errLengthPass
	}

	//hash password
	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("masalah hashing password")
	}
	user.Password = hashPass

	userUpdate, err := s.repo.UpdateUser(c, user)
	if err != nil {
		return &entity.User{}, err
	}
	//update user
	return userUpdate, nil
}

func (s *UserService) DeleteUser(c context.Context, id int) error {
	return s.repo.DeleteUser(c, id)
}

// Login user
func (s *UserService) LoginUser(c context.Context, email string, password string) (*entity.User, string, error) {

	//validasi data kosong
	errEmpty := validator.CheckDataEmpty(email, password)
	if errEmpty != nil {
		return nil, "", errEmpty
	}

	//validasi email
	errEmail := validator.EmailFormat(email)
	if errEmail != nil {
		fmt.Println("salah di emailformat")
		return nil, "", errEmail
	}

	user, err := s.repo.GetUserByEmail(c, email)
	if err != nil {
		fmt.Println("salah di getemail")
		return nil, "", err
	}

	if user == nil {
		return nil, "", errors.New("email tidak terdaftar")
	}

	//validasi password
	ComparePas := helper.CompareHash(user.Password, password)
	if !ComparePas {
		return nil, "", errors.New("password tidak sesuai")
	}

	token, err := common.GenerateAccessToken(c, user)
	if err != nil {
		fmt.Println("salah di handler generate access token")
		return nil, "", err
	}

	return user, token, nil
}

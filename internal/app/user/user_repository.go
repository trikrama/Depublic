package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface{
	GetAllUser(c context.Context)([]*User, error)
	GetUserByID(c context.Context, id int)(*User,error)
	CreateUser(c context.Context, user *User)error
	UpdateUser(c context.Context, user *User)error
	DeleteUser(c context.Context, id int)error
	GetUserByEmail(c context.Context, email string)(*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository)GetAllUser(c context.Context)([]*User, error){
	users := make([]*User, 0)
	err := r.db.WithContext(c).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository)GetUserByID(c context.Context, id int)(*User,error){
	user := new(User)
	err := r.db.WithContext(c).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository)CreateUser(c context.Context, user *User)error{
	err := r.db.WithContext(c).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository)UpdateUser(c context.Context, user *User)error{
	err := r.db.WithContext(c).Model(&User{}).Where("id = ?",user.ID).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository)DeleteUser(c context.Context, id int)error{
	err := r.db.WithContext(c).Delete(&User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository)GetUserByEmail(c context.Context, email string)(*User, error){
	user := new(User)
	err := r.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}



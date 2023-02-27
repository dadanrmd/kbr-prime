package userRepository

import (
	"kbrprime-be/internal/app/model/userModel"
)

type IUserRepository interface {
	FindAllUser() (*[]userModel.User, error)
	FindUserByID(id int64) (*userModel.User, error)
	FindUserByEmail(email string) (userData *userModel.User, err error)
	InsertUser(userData userModel.User) (*userModel.User, error)
	UpdateUser(userData userModel.User) (*userModel.User, error)
	FindUserByPhone(phone string) (userData *userModel.User, err error)
}

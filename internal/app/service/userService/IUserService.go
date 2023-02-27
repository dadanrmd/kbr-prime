package userService

import (
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/model/userModel"
)

type IUserService interface {
	FindAllUser() (data *[]userModel.User, err error)
	CreateNewUser(record *loggers.Data, req userModel.CreateUserReq) (*userModel.User, int, error)
	FindUserByEmail(email string) (*userModel.User, int, error)
	UpdateUser(request userModel.UpdateUserReq, userID int64) (*userModel.User, error)
	GetDetailUser(id int64) (data *userModel.User, err error)
}

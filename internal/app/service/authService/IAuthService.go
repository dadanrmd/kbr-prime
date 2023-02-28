package authService

import (
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/model/authModel"
)

type IAuthService interface {
	Login(record *loggers.Data, loginReq authModel.LoginReq) (loginRes authModel.LoginRes, err error)
	RevokeToken(token string) error
}

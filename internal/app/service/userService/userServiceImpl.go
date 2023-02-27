package userService

import (
	"errors"
	"fmt"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/symmetricHash"
	"kbrprime-be/internal/app/model/userModel"
	"kbrprime-be/internal/app/repository/userRepository"
	"net/http"

	"gorm.io/gorm"

	"github.com/rs/zerolog/log"
)

var (
	ErrEmailAlreadyExist    = errors.New("email sudah terdaftar")
	ErrUserNotFound         = errors.New("user tidak ditemukan")
	ErrorDuplicatePhone     = errors.New("nomor telepon sudah terdaftar")
	ErrorInvalidFormatEmail = errors.New("format email tidak valid")
	ErrorInvalidFormatPhone = errors.New("format nomor telepon tidak valid")
)

type userService struct {
	userRepo userRepository.IUserRepository
}

func NewUserService(userRepo userRepository.IUserRepository) IUserService {
	return &userService{userRepo}
}

func (u userService) FindUserByEmail(email string) (*userModel.User, int, error) {
	user, err := u.userRepo.FindUserByEmail(email)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return user, http.StatusInternalServerError, err
	} else {
		return user, http.StatusNotFound, err
	}
}

func (u userService) CreateNewUser(record *loggers.Data, req userModel.CreateUserReq) (*userModel.User, int, error) {
	_, err := u.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u.AddNewUser(record, req)
		}
	}
	return nil, http.StatusInternalServerError, err
}

func (u userService) AddNewUser(record *loggers.Data, req userModel.CreateUserReq) (*userModel.User, int, error) {
	subroutineName := "userService.AddNewUser"

	if req.NoHp != "" {
		_, err := u.userRepo.FindUserByPhone(req.NoHp)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				loggers.Log(record, fmt.Sprintf("%s, failed FindUserByPhone, err: %s", subroutineName, err.Error()))
				return nil, http.StatusInternalServerError, err
			}
		} else {
			loggers.Log(record, fmt.Sprintf("%s, failed ErrorDuplicatePhone, err: %s", subroutineName, ErrorDuplicatePhone.Error()))
			return nil, http.StatusBadRequest, ErrorDuplicatePhone
		}
	}

	createdUser, err := u.userRepo.InsertUser(userModel.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: symmetricHash.ToBcrypt(req.Password),
		NoHp:     req.NoHp,
		Role:     "USER",
		Status:   "AKTIF",
	})

	if err != nil {
		loggers.Log(record, fmt.Sprintf("%s, failed InsertUser, err: %s", subroutineName, err.Error()))
		return nil, http.StatusInternalServerError, err
	}

	return createdUser, http.StatusOK, nil
}

func (u userService) GetDetailUser(id int64) (data *userModel.User, err error) {
	data, err = u.userRepo.FindUserByID(id)
	if err != nil {
		err = ErrUserNotFound
		return
	}
	return
}

func (u userService) FindAllUser() (data *[]userModel.User, err error) {

	data, err = u.userRepo.FindAllUser()
	if err != nil {
		err = ErrUserNotFound
		return
	}
	return
}

func (u userService) UpdateUser(request userModel.UpdateUserReq, userID int64) (*userModel.User, error) {
	subroutineName := "userService.UpdateUser"

	existingUser, err := u.userRepo.FindUserByID(request.IdUser)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("%s, failed FindUserByID, err: %s", subroutineName, err.Error()))
		return nil, ErrUserNotFound
	}

	user, err := u.userRepo.FindUserByEmail(request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Msg(fmt.Sprintf("%s, failed FindUserByEmail, err: %s", subroutineName, err.Error()))
		return nil, err
	}

	if user.Email != "" && err == nil {
		log.Error().Msg(fmt.Sprintf("%s, failed ErrEmailAlreadyExist, err: %s", subroutineName, ErrEmailAlreadyExist.Error()))
		return nil, ErrEmailAlreadyExist
	}

	if request.NoHp != "" {
		user, err = u.userRepo.FindUserByPhone(request.NoHp)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Msg(fmt.Sprintf("%s, failed FindUserByPhone, err: %s", subroutineName, err.Error()))
			return nil, err
		}

		if user.NoHp != "" && err == nil {
			log.Error().Msg(fmt.Sprintf("%s, failed ErrorDuplicatePhone, err: %s", subroutineName, ErrorDuplicatePhone.Error()))
			return nil, ErrorDuplicatePhone
		}
	}
	//update data
	existingUser.Nama = request.Nama
	existingUser.Email = request.Email
	existingUser.NoHp = request.NoHp
	existingUser.Status = request.Status

	//save existing user data
	updatedUser, err := u.userRepo.UpdateUser(*existingUser)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("%s, failed UpdateUser, err: %s", subroutineName, err.Error()))
		return nil, err
	}

	return updatedUser, nil
}

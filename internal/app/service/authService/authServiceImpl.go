package authService

import (
	"errors"
	"kbrprime-be/internal/app/commons/jwtHelper"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/symmetricHash"
	"kbrprime-be/internal/app/middleware/authMiddleware"
	"kbrprime-be/internal/app/model/authModel"
	"kbrprime-be/internal/app/repository/userRepository"
	"os"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredential      = errors.New("invalid credential")
	ErrInvalidOTP             = errors.New("invalid otp")
	ErrUserNotFound           = errors.New("user not found")
	ErrExpiredOTP             = errors.New("otp expired")
	ErrUsedOTP                = errors.New("otp used")
	ErrRoleNotFound           = errors.New("role not found")
	ErrCannotSendNotification = errors.New("cannot send otp")
	ErrProductNotFound        = errors.New("product not found")
	ErrParseEmail             = errors.New("error when parsing email template")
)

const (
	MaxOTPVerificationLifetime = 5 * time.Minute
	OTPLength                  = 6
)

type authUseCase struct {
	userRepo userRepository.IUserRepository
}

func NewAuthService(userRepo userRepository.IUserRepository) IAuthService {
	return &authUseCase{userRepo}
}

func (a authUseCase) Login(record *loggers.Data, loginReq authModel.LoginReq) (loginRes authModel.LoginRes, err error) {
	userData, err := a.userRepo.FindUserByEmail(loginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userData, err = a.userRepo.FindUserByPhone(loginReq.Email)
			if err != nil {
				return loginRes, ErrInvalidCredential
			}
		} else {
			return loginRes, ErrInvalidCredential
		}
	}

	if !symmetricHash.CompareBcrypt(userData.Password, loginReq.Password) {
		return loginRes, ErrInvalidCredential
	}

	loginRes.User = authModel.UserLogin{
		IdUser:     userData.IdUser,
		Nik:        userData.Nik,
		Nama:       userData.Nama,
		Email:      userData.Email,
		Password:   userData.Password,
		NoHp:       userData.NoHp,
		Role:       userData.Role,
		Status:     userData.Status,
		Update_Key: userData.Update_Key,
	}

	jwtExpirationDurationDayString := os.Getenv("JWT_EXPIRATION_DURATION_DAY")
	var jwtExpirationDurationDay int
	jwtExpirationDurationDay, err = strconv.Atoi(jwtExpirationDurationDayString)
	if err != nil {
		return loginRes, err
	}

	// Conversion to seconds
	jwtExpiredAt := time.Now().Unix() + int64(jwtExpirationDurationDay*3600*24)
	tokenUID := uuid.NewV4().String() + "00" + cast.ToString(userData.IdUser)

	userClaims := jwtHelper.CustomClaims{IdUser: userData.IdUser, ExpiresAt: jwtExpiredAt, TokenUID: tokenUID}
	jwtToken, err := jwtHelper.NewWithClaims(userClaims)
	if err != nil {
		return loginRes, err
	}

	loginRes.Token = jwtToken
	return loginRes, nil
}

func (a authUseCase) RevokeToken(token string) error {
	//extract token claims
	//Extract JWT Token from Bearer
	jwtTokenSplit := strings.Split(token, "Bearer ")
	if jwtTokenSplit[1] == "" {
		return authMiddleware.ErrInvalidToken
	}
	tokenSegment := jwtTokenSplit[1]

	_, err := jwtHelper.ExtractClaims(tokenSegment)
	if err != nil {
		return authMiddleware.ErrInvalidToken
	}

	return nil
}

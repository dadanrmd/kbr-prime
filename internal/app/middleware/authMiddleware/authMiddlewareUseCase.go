package authMiddleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/jwtHelper"
	"kbrprime-be/internal/app/model/userModel"
	"kbrprime-be/internal/app/repository/userRepository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrUserNotFound = errors.New("user not found")
	ErrURLNotFound  = errors.New("url token not found")
	ErrTokenRevoked = errors.New("token revoked")
)

type authMiddleware struct {
	userRepo userRepository.IUserRepository
}

func NewAuthMiddleware(userRepo userRepository.IUserRepository) IAuthMiddleware {
	return &authMiddleware{userRepo}
}

func (auth *authMiddleware) AuthorizeUserRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		bearerTokenSplit := strings.Split(bearerToken, " ")

		if len(bearerTokenSplit) < 2 {
			res := jsonHttpResponse.FailedResponse{
				Status:       jsonHttpResponse.FailedStatus,
				ResponseCode: "00",
				Message:      "invalid token",
			}
			jsonHttpResponse.Unauthorized(c, res)
			c.Abort()
			return
		}

		jwtToken := bearerTokenSplit[1]
		userData, err := auth.getUserFromJWTWithRoleValidation(jwtToken)
		if err != nil {
			if err == jwtHelper.ErrTokenExpired {
				res := jsonHttpResponse.FailedResponse{
					Status:       jsonHttpResponse.FailedStatus,
					ResponseCode: "00",
					Message:      err.Error(),
				}
				jsonHttpResponse.Unauthorized(c, res)
				c.Abort()
				return
			}

			res := jsonHttpResponse.FailedResponse{
				Status:       jsonHttpResponse.FailedStatus,
				ResponseCode: "00",
				Message:      "invalid token",
			}
			jsonHttpResponse.InternalServerError(c, res)
			c.Abort()
			return
		}

		//put into user context, convert user refresh token claims to user claim
		prospectClaims := jwtHelper.CustomClaims{
			IdUser: userData.IdUser,
		}

		c.Set("user", prospectClaims)
	}
}
func (auth *authMiddleware) AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		bearerTokenSplit := strings.Split(bearerToken, " ")

		if len(bearerTokenSplit) < 2 {
			res := jsonHttpResponse.FailedResponse{
				Status:       jsonHttpResponse.FailedStatus,
				ResponseCode: "00",
				Message:      "invalid token",
			}
			jsonHttpResponse.Unauthorized(c, res)
			c.Abort()
			return
		}

		jwtToken := bearerTokenSplit[1]
		userData, err := auth.getUserFromJWTWithRoleValidation(jwtToken)
		if err != nil {
			if err == jwtHelper.ErrTokenExpired {
				res := jsonHttpResponse.FailedResponse{
					Status:       jsonHttpResponse.FailedStatus,
					ResponseCode: "00",
					Message:      err.Error(),
				}
				jsonHttpResponse.Unauthorized(c, res)
				c.Abort()
				return
			}

			res := jsonHttpResponse.FailedResponse{
				Status:       jsonHttpResponse.FailedStatus,
				ResponseCode: "00",
				Message:      "invalid token",
			}
			jsonHttpResponse.InternalServerError(c, res)
			c.Abort()
			return
		}

		test, _ := json.Marshal(userData)
		fmt.Println("userData->", cast.ToString(test))
		//put into user context
		c.Set("user", userData)
	}
}

func (auth *authMiddleware) getUserFromJWTWithRoleValidation(jwtToken string) (user userModel.User, err error) {
	if jwtToken == "" {
		return user, ErrInvalidToken
	}
	jwtTokenClaims, err := jwtHelper.VerifyTokenWithClaims(jwtToken)
	if err != nil {
		return user, ErrInvalidToken
	}
	userData, err := auth.userRepo.FindUserByID(jwtTokenClaims.IdUser)
	if err != nil {
		return user, ErrUserNotFound
	}

	return *userData, nil
}

package RouteHelpers

import (
	"errors"
	"kbrprime-be/internal/app/model/userModel"

	"github.com/gin-gonic/gin"
)

var (
	ErrorContextNotExist     = errors.New("user context not exist")
	ErrorParsingUserModel    = errors.New("error parsing user model")
	ErrInvalidJWTForProspect = errors.New("invalid jwt, forbidden access")
)

func GetUserFromJWTContext(c *gin.Context) (*userModel.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, ErrorContextNotExist
	}

	if user == "" {
		return nil, ErrorContextNotExist
	}

	userData, ok := user.(*userModel.User)
	if !ok {
		return nil, ErrorParsingUserModel
	}

	return userData, nil
}

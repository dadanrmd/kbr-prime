package authMiddleware

import (
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	AuthorizeUser() gin.HandlerFunc
	AuthorizeUserRefreshToken() gin.HandlerFunc
}

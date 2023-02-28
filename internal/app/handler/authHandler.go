package handler

import (
	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/requestvalidationerror"
	"kbrprime-be/internal/app/commons/utils"
	"kbrprime-be/internal/app/model/authModel"
	"kbrprime-be/internal/app/service/authService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	HandlerOption
}

func (authDelivery AuthHandler) Login(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request authModel.LoginReq
	errBind := c.ShouldBind(&request)
	if errBind != nil {
		validations := requestvalidationerror.GetvalidationError(errBind)

		if len(validations) > 0 {
			jsonHttpResponse.NewFailedMissingRequiredFieldResponse(c, validations)
			return
		}
		utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, errBind.Error(), "")
		return
	}

	loginRes, err := authDelivery.AuthService.Login(record, request)
	if err != nil {
		if err == authService.ErrInvalidCredential {
			utils.BasicResponse(record, c.Writer, false, http.StatusUnauthorized, err.Error(), "")
			return
		}
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, loginRes, "Success")
}

func (authDelivery AuthHandler) RevokeToken(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	err := authDelivery.AuthService.RevokeToken(c.GetHeader("Authorization"))
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "token revoked successfully")
}

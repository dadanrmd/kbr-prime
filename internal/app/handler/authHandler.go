package handler

import (
	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/requestvalidationerror"
	"kbrprime-be/internal/app/model/authModel"
	"kbrprime-be/internal/app/service/authService"

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

		jsonHttpResponse.NewFailedBadRequestResponse(c, errBind.Error())
		return
	}

	loginRes, err := authDelivery.AuthService.Login(record, request)
	if err != nil {
		if err == authService.ErrInvalidCredential {
			errPayload := jsonHttpResponse.FailedResponse{
				Status:       jsonHttpResponse.FailedStatus,
				ResponseCode: "00",
				Message:      err.Error(),
			}
			jsonHttpResponse.Unauthorized(c, errPayload)
			return
		}
		errPayload := jsonHttpResponse.FailedResponse{
			Status:       jsonHttpResponse.FailedStatus,
			ResponseCode: "00",
			Message:      err.Error(),
		}
		jsonHttpResponse.InternalServerError(c, errPayload)
		return
	}

	successPayload := jsonHttpResponse.SuccessResponse{
		Status:       jsonHttpResponse.SuccessStatus,
		ResponseCode: "00",
		Message:      "",
		Data:         loginRes,
	}
	jsonHttpResponse.OK(c, successPayload)
	return
}

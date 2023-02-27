package handler

import (
	"kbrprime-be/internal/app/commons/RouteHelpers"
	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/requestvalidationerror"
	"kbrprime-be/internal/app/commons/utils"
	"kbrprime-be/internal/app/model/userModel"
	"kbrprime-be/internal/app/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserHandler struct {
	HandlerOption
}

func (userDelivery UserHandler) GetAllUser(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAllUserRes, err := userDelivery.UserService.FindAllUser()
	if err != nil {
		if err == userService.ErrEmailAlreadyExist {
			utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "")
			return
		}
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAllUserRes, "displayed all user data")
}

func (userDelivery UserHandler) UpdateUser(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request userModel.UpdateUserReq
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

	if request.NoHp != "" {
		if !requestvalidationerror.IsINAPhoneValid(request.NoHp) {
			jsonHttpResponse.NewFailedBadRequestResponse(c, userService.ErrorInvalidFormatPhone.Error())
			return
		}
	}

	if !requestvalidationerror.IsEmailValid(request.Email) {
		jsonHttpResponse.NewFailedBadRequestResponse(c, userService.ErrorInvalidFormatEmail.Error())
		return
	}

	requestedUser, err := RouteHelpers.GetUserFromJWTContext(c)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}

	createNewUserRes, err := userDelivery.UserService.UpdateUser(request, requestedUser.IdUser)
	if err != nil {
		if err == userService.ErrEmailAlreadyExist {
			utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "")
			return
		}
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, createNewUserRes, userModel.MsgSuccessUpdateUser)
}

func (userDelivery UserHandler) AddNewUser(c *gin.Context) {
	record := loggers.StartRecord(c.Request)

	var request userModel.CreateUserReq
	errBind := c.ShouldBind(&request)
	if errBind != nil {
		validations := requestvalidationerror.GetvalidationError(errBind)

		if len(validations) > 0 {
			loggers.EndRecord(record, errBind.Error(), http.StatusBadRequest)
			jsonHttpResponse.NewFailedMissingRequiredFieldResponse(c, validations)
			return
		}

		utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, userModel.MsgFailedAddUser, errBind.Error())
		return
	}

	if request.NoHp != "" {
		if !requestvalidationerror.IsINAPhoneValid(request.NoHp) {
			utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, userModel.MsgFailedAddUser, userService.ErrorInvalidFormatPhone.Error())
			return
		}
	}

	if !requestvalidationerror.IsEmailValid(request.Email) {
		utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, userModel.MsgFailedAddUser, userService.ErrorInvalidFormatEmail.Error())
		return
	}

	user, statusCode, err := userDelivery.UserService.CreateNewUser(record, request)
	if err != nil || statusCode == http.StatusInternalServerError {
		utils.BasicResponse(record, c.Writer, false, statusCode, userModel.MsgFailedAddUser, "")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, user, "Registrasi berhasil")
}

func (userDelivery UserHandler) GetDetailUser(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	id := c.Param("id")
	user, err := userDelivery.UserService.GetDetailUser(cast.ToInt64(id))
	if err != nil {
		if err == userService.ErrEmailAlreadyExist {
			utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "")
			return
		}
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, user, "displayed user data")
}

func (userDelivery UserHandler) GetUserByEmail(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	id := c.Param("id")
	user, httpStatus, err := userDelivery.UserService.FindUserByEmail(id)
	if err != nil {
		if httpStatus == http.StatusNotFound {
			utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "")
		} else {
			utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, user, "displayed user data")
}

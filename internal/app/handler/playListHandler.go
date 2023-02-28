package handler

import (
	"net/http"

	"kbrprime-be/internal/app/commons/RouteHelpers"
	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/requestvalidationerror"
	"kbrprime-be/internal/app/model/playListModel"

	"kbrprime-be/internal/app/commons/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type PlayListHandler struct {
	HandlerOption
}

func (handler PlayListHandler) GetPlayList(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	requestedUser, err := RouteHelpers.GetUserFromJWTContext(c)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusUnauthorized, err.Error(), "")
		return
	}

	getAll, err := handler.PlayListService.ListPlayList(record, requestedUser.IdUser)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler PlayListHandler) AddPlayList(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request playListModel.ReqPlayList
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

	requestedUser, err := RouteHelpers.GetUserFromJWTContext(c)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	request.UserId = requestedUser.IdUser

	getAll, err := handler.PlayListService.AddPlayList(record, request)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler PlayListHandler) DeletePlayList(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	id := c.Param("id")
	err := handler.PlayListService.DeletePlayList(record, cast.ToInt64(id))
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success")
}

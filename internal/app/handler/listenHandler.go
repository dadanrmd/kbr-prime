package handler

import (
	"net/http"

	"kbrprime-be/internal/app/commons/jsonHttpResponse"
	"kbrprime-be/internal/app/commons/loggers"
	"kbrprime-be/internal/app/commons/requestvalidationerror"
	"kbrprime-be/internal/app/model/listenModel"
	"kbrprime-be/internal/app/model/userModel"

	"kbrprime-be/internal/app/commons/utils"

	"github.com/gin-gonic/gin"
)

type ListenHandler struct {
	HandlerOption
}

func (handler ListenHandler) RecordData(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request listenModel.ReqRecord
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
	data, err := handler.ListenService.RecordEpisode(record, request)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, data, "Success")
}

package handler

import (
	"net/http"

	"kbrprime-be/internal/app/commons/loggers"

	"kbrprime-be/internal/app/commons/utils"

	"github.com/gin-gonic/gin"
)

type HealthyCheckHandler struct {
	HandlerOption
}

func (handler HealthyCheckHandler) HealthyCheck(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.HealtyService.FindAllHealty(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Healthy check")
}

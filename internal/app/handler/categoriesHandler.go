package handler

import (
	"net/http"

	"kbrprime-be/internal/app/commons/loggers"

	"kbrprime-be/internal/app/commons/utils"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	HandlerOption
}

func (handler CategoriesHandler) GetAll(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.CategoriesService.FindAllCategories(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

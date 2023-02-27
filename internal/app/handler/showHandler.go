package handler

import (
	"net/http"

	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/commons/loggers"

	"kbrprime-be/internal/app/commons/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ShowHandler struct {
	HandlerOption
}

func (handler ShowHandler) GetAll(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.ShowService.FindAllShow(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler ShowHandler) LatestNews(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.ShowService.LatestNews(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler ShowHandler) LatestEpisodes(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.ShowService.LatestEpisodes(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler ShowHandler) ListNews(c *gin.Context) {
	record := loggers.StartRecord(c.Request)

	pageNumber := cast.ToInt(c.Query("page_number"))
	pageSize := cast.ToInt(c.Query("page_size"))

	paging := datapaging.Datapaging{Page: pageNumber, Limit: pageSize}

	res, count, err := handler.ShowService.GetNewsWithPaging(record, paging)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "Data tidak ditemukan")
		return
	}
	data := map[string]interface{}{
		"page_number":        pageNumber,
		"page_size":          pageSize,
		"total_record_count": count,
		"records":            res,
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, data, "successfully")
}

func (handler ShowHandler) ListEpisodes(c *gin.Context) {
	record := loggers.StartRecord(c.Request)

	pageNumber := cast.ToInt(c.Query("page_number"))
	pageSize := cast.ToInt(c.Query("page_size"))

	paging := datapaging.Datapaging{Page: pageNumber, Limit: pageSize}

	res, count, err := handler.ShowService.GetEpisodesWithPaging(record, paging)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "Data tidak ditemukan")
		return
	}
	data := map[string]interface{}{
		"page_number":        pageNumber,
		"page_size":          pageSize,
		"total_record_count": count,
		"records":            res,
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, data, "successfully")
}

func (handler ShowHandler) TopThree(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.ShowService.GetTopThree(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

func (handler ShowHandler) Sorotan(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	getAll, err := handler.ShowService.GetSorotan(record)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "")
		return
	}
	utils.BasicResponse(record, c.Writer, true, http.StatusOK, getAll, "Success")
}

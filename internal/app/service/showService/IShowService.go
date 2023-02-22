package showService

import (
	"kbrprime-be/internal/app/model/showModel"

	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/commons/loggers"
)

type IShowService interface {
	FindAllShow(record *loggers.Data) (data *[]showModel.Show, err error)
	LatestNews(record *loggers.Data) (data *[]showModel.Show, err error)
	LatestEpisodes(record *loggers.Data) (data *[]showModel.Show, err error)
	GetNewsWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error)
	GetEpisodesWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error)
}

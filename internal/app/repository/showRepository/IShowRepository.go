package showRepository

import (
	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/model/showModel"
)

type IShowRepository interface {
	FindAll() (data *[]showModel.Show, err error)
	GetLatestNews(categoryID []int64) (data *[]showModel.Show, err error)
	GetLatestEpisodes(categoryID []int64) (data *[]showModel.Show, err error)
	GetNewsByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]showModel.Show, count int64, err error)
	GetEpisodesByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]showModel.Show, count int64, err error)
}

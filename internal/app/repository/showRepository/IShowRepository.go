package showRepository

import (
	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/model/showModel"
)

type IShowRepository interface {
	FindAll() (data *[]showModel.Show, err error)
	GetLatestNews() (data *[]showModel.Show, err error)
	GetLatestEpisodes() (data *[]showModel.Show, err error)
	GetNewsByStatusTags(paging datapaging.Datapaging, tag string) (data *[]showModel.Show, count int64, err error)
	GetEpisodesByStatusTags(paging datapaging.Datapaging, tag string) (data *[]showModel.Show, count int64, err error)
}

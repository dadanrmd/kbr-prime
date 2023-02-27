package showRepository

import (
	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/model/podcastModel"
	"kbrprime-be/internal/app/model/showModel"
)

type IShowRepository interface {
	FindAll() (data *[]showModel.Show, err error)
	GetLatestNews(categoryID []int64) (data *[]podcastModel.Podcast, err error)
	GetLatestEpisodes(categoryID []int64) (data *[]podcastModel.Podcast, err error)
	GetNewsByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]podcastModel.Podcast, count int64, err error)
	GetEpisodesByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]podcastModel.Podcast, count int64, err error)
	GetTopByLimit(limit, offset int) (data *[]podcastModel.Podcast, err error)
}

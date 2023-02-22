package showService

import (
	"errors"
	"kbrprime-be/internal/app/model/showModel"
	"kbrprime-be/internal/app/repository/showRepository"

	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errBerita  = errors.New("Berita tidak ada")
	errEpisode = errors.New("Episode tidak ada")
)

type showService struct {
	showRepo showRepository.IShowRepository
}

func NewShowService(show showRepository.IShowRepository) IShowService {
	return &showService{showRepo: show}
}

func (h showService) FindAllShow(record *loggers.Data) (data *[]showModel.Show, err error) {
	loggers.Logf(record, "Info, FindAllShow")
	data, err = h.showRepo.FindAll()
	if err != nil {
		loggers.Logf(record, "Error, FindAll")
		err = errBerita
		return
	}
	return
}

func (h showService) LatestNews(record *loggers.Data) (data *[]showModel.Show, err error) {
	loggers.Logf(record, "Info, LatestNews")
	data, err = h.showRepo.GetLatestNews()
	if err != nil {
		loggers.Logf(record, "Error, GetLatestNews")
		err = errBerita
		return
	}
	return
}

func (h showService) LatestEpisodes(record *loggers.Data) (data *[]showModel.Show, err error) {
	loggers.Logf(record, "Info, LatestNews")
	data, err = h.showRepo.GetLatestEpisodes()
	if err != nil {
		loggers.Logf(record, "Error, GetLatestEpisodes")
		err = errBerita
		return
	}
	return
}

func (h showService) GetNewsWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error) {
	loggers.Logf(record, "Info, GetNewsWithPaging")
	data, count, err = h.showRepo.GetNewsByStatusTags(paging, "news")
	if err != nil {
		err = errBerita
		loggers.Logf(record, "Error, GetNewsByStatusTags")
		return
	}
	return
}

func (h showService) GetEpisodesWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error) {
	loggers.Logf(record, "Info, GetEpisodesWithPaging")
	data, count, err = h.showRepo.GetEpisodesByStatusTags(paging, "news")
	if err != nil {
		err = errEpisode
		loggers.Logf(record, "Error, GetEpisodesByStatusTags")
		return
	}
	return
}

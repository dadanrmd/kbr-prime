package showService

import (
	"errors"
	"kbrprime-be/internal/app/model/showModel"
	"kbrprime-be/internal/app/repository/categoriesRepository"
	"kbrprime-be/internal/app/repository/showRepository"

	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errBerita  = errors.New("Berita tidak ada")
	errEpisode = errors.New("Episode tidak ada")
)

type showService struct {
	showRepo       showRepository.IShowRepository
	categoriesRepo categoriesRepository.ICategoriesRepository
}

func NewShowService(showRepo showRepository.IShowRepository, categoriesRepo categoriesRepository.ICategoriesRepository) IShowService {
	return &showService{showRepo, categoriesRepo}
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
	var categoryID []int64
	category, err := h.categoriesRepo.FindByCategoryScope("News")
	if err != nil {
		loggers.Logf(record, "Error, FindByCategoryScope")
	}
	for _, v := range *category {
		categoryID = append(categoryID, v.IdCategory)
	}
	data, err = h.showRepo.GetLatestNews(categoryID)
	if err != nil {
		loggers.Logf(record, "Error, GetLatestNews")
		err = errBerita
		return
	}
	return
}

func (h showService) LatestEpisodes(record *loggers.Data) (data *[]showModel.Show, err error) {
	loggers.Logf(record, "Info, LatestEpisodes")
	var categoryID []int64
	category, err := h.categoriesRepo.FindByCategoryScope("News")
	if err != nil {
		loggers.Logf(record, "Error, FindByCategoryScope")
	}
	for _, v := range *category {
		categoryID = append(categoryID, v.IdCategory)
	}
	data, err = h.showRepo.GetLatestEpisodes(categoryID)
	if err != nil {
		loggers.Logf(record, "Error, GetLatestEpisodes")
		err = errBerita
		return
	}
	return
}

func (h showService) GetNewsWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error) {
	loggers.Logf(record, "Info, GetNewsWithPaging")
	var categoryID []int64
	category, err := h.categoriesRepo.FindByCategoryScope("News")
	if err != nil {
		loggers.Logf(record, "Error, FindByCategoryScope")
	}
	for _, v := range *category {
		categoryID = append(categoryID, v.IdCategory)
	}
	data, count, err = h.showRepo.GetNewsByStatusTags(paging, categoryID)
	if err != nil {
		err = errBerita
		loggers.Logf(record, "Error, GetNewsByStatusTags")
		return
	}
	return
}

func (h showService) GetEpisodesWithPaging(record *loggers.Data, paging datapaging.Datapaging) (data *[]showModel.Show, count int64, err error) {
	loggers.Logf(record, "Info, GetEpisodesWithPaging")
	var categoryID []int64
	category, err := h.categoriesRepo.FindByCategoryScope("News")
	if err != nil {
		loggers.Logf(record, "Error, FindByCategoryScope")
	}
	for _, v := range *category {
		categoryID = append(categoryID, v.IdCategory)
	}
	data, count, err = h.showRepo.GetEpisodesByStatusTags(paging, categoryID)
	if err != nil {
		err = errEpisode
		loggers.Logf(record, "Error, GetEpisodesByStatusTags")
		return
	}
	return
}

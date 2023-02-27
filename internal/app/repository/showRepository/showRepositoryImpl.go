package showRepository

import (
	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/model/showModel"

	"gorm.io/gorm"
)

type showRepository struct {
	db *gorm.DB
}

func NewShowRepository(db *gorm.DB) IShowRepository {
	return &showRepository{db: db}
}
func (h showRepository) FindAll() (data *[]showModel.Show, err error) {
	err = h.db.Find(&data).Error
	return
}

func (h showRepository) GetLatestNews(categoryID []int64) (data *[]showModel.Show, err error) {
	err = h.db.Where("category in ? and date_created < curdate() - INTERVAL DAYOFWEEK(curdate()) DAY and status = ?", categoryID, "AKTIF").Limit(4).Order("date_created desc").Find(&data).Error
	return
}

func (h showRepository) GetLatestEpisodes(categoryID []int64) (data *[]showModel.Show, err error) {
	err = h.db.Where("category not in ? and date_created < curdate() - INTERVAL 1 DAY and status = ?", categoryID, "AKTIF").Limit(4).Order("date_created desc").Find(&data).Error
	return
}

func (h showRepository) GetNewsByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]showModel.Show, count int64, err error) {
	db := h.db.Model(&data).Where("category in ?", categoryID)
	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	err = db.Order("date_created desc").Find(&data).Error

	return data, count, err
}

func (h showRepository) GetEpisodesByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]showModel.Show, count int64, err error) {
	db := h.db.Model(&data).Where("category not in ?", categoryID)
	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	err = db.Order("date_created desc").Find(&data).Error

	return data, count, err
}

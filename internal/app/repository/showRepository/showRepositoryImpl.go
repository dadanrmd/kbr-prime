package showRepository

import (
	datapaging "kbrprime-be/internal/app/commons/dataPagingHelper"
	"kbrprime-be/internal/app/model/podcastModel"
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

func (h showRepository) GetLatestNews(categoryID []int64) (data *[]podcastModel.Podcast, err error) {
	db := h.db.Model(&podcastModel.Podcast{})
	db.Joins("JOIN kbr_shows on kbr_shows.id_show = kbr_podcasts.id_show")
	err = db.Where("kbr_shows.category in ? and kbr_podcasts.publish_date < curdate() - INTERVAL DAYOFWEEK(curdate()) DAY and kbr_shows.status = ?", categoryID, "AKTIF").Limit(4).Order("kbr_podcasts.publish_date desc").Find(&data).Error
	return
}

func (h showRepository) GetLatestEpisodes(categoryID []int64) (data *[]podcastModel.Podcast, err error) {
	db := h.db.Model(&podcastModel.Podcast{})
	db.Joins("JOIN kbr_shows on kbr_shows.id_show = kbr_podcasts.id_show")
	err = db.Where("kbr_shows.category not in ? and kbr_podcasts.publish_date < curdate() - INTERVAL 1 DAY and kbr_shows.status = ?", categoryID, "AKTIF").Limit(4).Order("kbr_podcasts.publish_date desc").Find(&data).Error
	return
}

func (h showRepository) GetNewsByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]podcastModel.Podcast, count int64, err error) {
	db := h.db.Model(&podcastModel.Podcast{}).Joins("JOIN kbr_shows on kbr_shows.id_show = kbr_podcasts.id_show").Where("kbr_shows.category in ?", categoryID)
	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	err = db.Order("kbr_podcasts.publish_date desc").Find(&data).Error

	return data, count, err
}

func (h showRepository) GetEpisodesByStatusTags(paging datapaging.Datapaging, categoryID []int64) (data *[]podcastModel.Podcast, count int64, err error) {
	db := h.db.Model(&podcastModel.Podcast{}).Joins("JOIN kbr_shows on kbr_shows.id_show = kbr_podcasts.id_show").Where("kbr_shows.category not in ?", categoryID)
	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	err = db.Order("kbr_podcasts.publish_date desc").Find(&data).Error

	return data, count, err
}

package playListRepository

import (
	"kbrprime-be/internal/app/model/playListModel"
	"kbrprime-be/internal/app/model/podcastModel"

	"gorm.io/gorm"
)

type playListRepository struct {
	db *gorm.DB
}

func NewPlayListRepository(db *gorm.DB) IPlayListRepository {
	return &playListRepository{db: db}
}

func (p playListRepository) FindByUser(UserId int64) (data *[]playListModel.ListPlayList, err error) {
	db := p.db.Model(&podcastModel.Podcast{}).Select("kbr_podcasts.*, kbr_play_list.id as id_play_list")
	db.Joins("JOIN kbr_play_list on kbr_play_list.episode_id = kbr_podcasts.id_podcast")
	err = db.Where("kbr_play_list.user_id = ?", UserId).Order("kbr_play_list.created_at desc").Find(&data).Error
	return
}

func (p playListRepository) FindById(id int64) (data *playListModel.PlayList, err error) {
	err = p.db.Where("id = ?", id).First(&data).Error
	return
}

func (p playListRepository) Insert(data playListModel.PlayList) (*playListModel.PlayList, error) {
	db := p.db.Create(&data)
	return &data, db.Error
}

func (p playListRepository) Delete(data playListModel.PlayList) error {
	db := p.db.Delete(&data)
	return db.Error
}

package listenRepository

import (
	"kbrprime-be/internal/app/model/listenModel"

	"gorm.io/gorm"
)

type listenRepository struct {
	db *gorm.DB
}

func NewListenRepository(db *gorm.DB) IListenRepository {
	return &listenRepository{db: db}
}
func (l listenRepository) FindByEpisode() (data *[]listenModel.Listen, err error) {
	err = l.db.Find(&data).Error
	return
}

func (l listenRepository) FindByUser() (data *[]listenModel.Listen, err error) {
	err = l.db.Find(&data).Error
	return
}

func (l listenRepository) FindByEpisodeAndUser(episodeId, userId int64, sessionId string) (data *listenModel.Listen, err error) {
	err = l.db.Where("episode_id = ? and user_id = ? and session_id = ?", episodeId, userId, sessionId).Last(&data).Error
	return
}

func (l listenRepository) InsertRecord(data listenModel.Listen) (*listenModel.Listen, error) {
	db := l.db.Create(&data)
	return &data, db.Error
}

func (l listenRepository) UpdateRecord(data listenModel.Listen) (*listenModel.Listen, error) {
	db := l.db.Save(&data)
	return &data, db.Error
}

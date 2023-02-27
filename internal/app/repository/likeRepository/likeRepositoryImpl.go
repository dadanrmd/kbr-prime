package likeRepository

import (
	"kbrprime-be/internal/app/model/likeModel"

	"gorm.io/gorm"
)

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &likeRepository{db: db}
}

func (l likeRepository) FindByEpisodeAndUser(UserId, EpisodeId int64) (data *likeModel.Like, err error) {
	err = l.db.Where("user_id = ? and episode_id = ?", UserId, EpisodeId).Last(&data).Error
	return
}

func (l likeRepository) Insert(data likeModel.Like) (*likeModel.Like, error) {
	db := l.db.Create(&data)
	return &data, db.Error
}

func (l likeRepository) Update(data likeModel.Like) (*likeModel.Like, error) {
	db := l.db.Save(&data)
	return &data, db.Error
}

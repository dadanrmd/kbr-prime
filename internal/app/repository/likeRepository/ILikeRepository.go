package likeRepository

import (
	"kbrprime-be/internal/app/model/likeModel"
)

type ILikeRepository interface {
	FindByEpisodeAndUser(UserId, EpisodeId int64) (data *likeModel.Like, err error)
	Insert(data likeModel.Like) (*likeModel.Like, error)
	Update(data likeModel.Like) (*likeModel.Like, error)
}

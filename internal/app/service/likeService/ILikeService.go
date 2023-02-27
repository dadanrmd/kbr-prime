package likeService

import (
	"kbrprime-be/internal/app/model/likeModel"

	"kbrprime-be/internal/app/commons/loggers"
)

type ILikeService interface {
	LikeEpisode(record *loggers.Data, req likeModel.ReqLike) (data *likeModel.Like, err error)
}

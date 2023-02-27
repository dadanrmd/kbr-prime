package likeService

import (
	"errors"
	"kbrprime-be/internal/app/model/likeModel"
	"kbrprime-be/internal/app/repository/likeRepository"
	"time"

	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errKategori = errors.New("Katgori tidak ada")
)

type likeService struct {
	likeRepo likeRepository.ILikeRepository
}

func NewLikeService(likeRepo likeRepository.ILikeRepository) ILikeService {
	return &likeService{likeRepo}
}

func (h likeService) LikeEpisode(record *loggers.Data, req likeModel.ReqLike) (data *likeModel.Like, err error) {
	loggers.Logf(record, "Info, Like")
	payload := likeModel.Like{
		UserId:    req.UserId,
		EpisodeId: req.EpisodeId,
	}
	dataEpisodeUser, err := h.likeRepo.FindByEpisodeAndUser(req.UserId, req.EpisodeId)
	if err != nil {
		loggers.Logf(record, "Error, FindByEpisodeAndUser")
		data, err = h.likeRepo.Insert(payload)
		if err != nil {
			loggers.Logf(record, "Error, Insert")
		}
	} else {
		if dataEpisodeUser.DeletedAt == 0 {
			dataEpisodeUser.DeletedAt = time.Now().UnixNano() / int64(time.Millisecond)
		} else {
			dataEpisodeUser.DeletedAt = 0
		}
		data, err = h.likeRepo.Update(*dataEpisodeUser)
		if err != nil {
			loggers.Logf(record, "Error, Delete")
		}
	}
	return
}

package playListService

import (
	"errors"
	"kbrprime-be/internal/app/model/playListModel"
	"kbrprime-be/internal/app/repository/playListRepository"

	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errKategori = errors.New("Katgori tidak ada")
)

type playListService struct {
	playListRepo playListRepository.IPlayListRepository
}

func NewPlayListService(playListRepo playListRepository.IPlayListRepository) IPlayListService {
	return &playListService{playListRepo}
}

func (h playListService) ListPlayList(record *loggers.Data, userId int64) (data *[]playListModel.ListPlayList, err error) {
	loggers.Logf(record, "Info, ListPlayList")
	data, err = h.playListRepo.FindByUser(userId)
	if err != nil {
		loggers.Logf(record, "Error, FindByUser")
	}
	return
}

func (h playListService) AddPlayList(record *loggers.Data, req playListModel.ReqPlayList) (data *playListModel.PlayList, err error) {
	loggers.Logf(record, "Info, AddPlayList")
	payload := playListModel.PlayList{
		UserId:    req.UserId,
		EpisodeId: req.EpisodeId,
	}
	data, err = h.playListRepo.Insert(payload)
	if err != nil {
		loggers.Logf(record, "Error, Insert")
	}
	return
}

func (h playListService) DeletePlayList(record *loggers.Data, id int64) (err error) {
	loggers.Logf(record, "Info, DeletePlayList")
	data, err := h.playListRepo.FindById(id)
	if err != nil {
		loggers.Logf(record, "Error, FindById")
		return
	}
	err = h.playListRepo.Delete(*data)
	if err != nil {
		loggers.Logf(record, "Error, Delete")
	}
	return
}

package listenService

import (
	"errors"
	"kbrprime-be/internal/app/model/listenModel"
	"kbrprime-be/internal/app/repository/listenRepository"

	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errKategori = errors.New("Katgori tidak ada")
)

type listenService struct {
	listenRepo listenRepository.IListenRepository
}

func NewListenService(listenRepo listenRepository.IListenRepository) IListenService {
	return &listenService{listenRepo}
}

func (h listenService) RecordEpisode(record *loggers.Data, req listenModel.ReqRecord) (data *listenModel.Listen, err error) {
	loggers.Logf(record, "Info, FindAllListen")
	dataEpisodeUser, err := h.listenRepo.FindByEpisodeAndUser(req.EpisodeId, req.UserId, req.SessionId)
	if err != nil {
		loggers.Logf(record, "Error, FindByEpisodeAndUser")
		payload := listenModel.Listen{
			UserId:    req.UserId,
			SessionId: req.SessionId,
			EpisodeId: req.EpisodeId,
			Play:      req.Play,
		}
		data, err = h.listenRepo.InsertRecord(payload)
		if err != nil {
			loggers.Logf(record, "Error, InsertRecord")
		}
	} else {
		data.Play = req.Play
		data, err = h.listenRepo.UpdateRecord(*dataEpisodeUser)
		if err != nil {
			loggers.Logf(record, "Error, InsertRecord")
		}
	}
	return
}

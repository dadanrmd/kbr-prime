package listenRepository

import (
	"kbrprime-be/internal/app/model/listenModel"
)

type IListenRepository interface {
	FindByEpisode() (data *[]listenModel.Listen, err error)
	FindByUser() (data *[]listenModel.Listen, err error)
	FindByEpisodeAndUser(episodeId, userId int64, sessionId string) (data *listenModel.Listen, err error)
	InsertRecord(data listenModel.Listen) (*listenModel.Listen, error)
	UpdateRecord(data listenModel.Listen) (*listenModel.Listen, error)
}

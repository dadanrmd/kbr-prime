package listenService

import (
	"kbrprime-be/internal/app/model/listenModel"

	"kbrprime-be/internal/app/commons/loggers"
)

type IListenService interface {
	RecordEpisode(record *loggers.Data, req listenModel.ReqRecord) (data *listenModel.Listen, err error)
}

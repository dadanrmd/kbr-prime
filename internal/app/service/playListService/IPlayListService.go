package playListService

import (
	"kbrprime-be/internal/app/model/playListModel"

	"kbrprime-be/internal/app/commons/loggers"
)

type IPlayListService interface {
	ListPlayList(record *loggers.Data, userId int64) (data *[]playListModel.ListPlayList, err error)
	AddPlayList(record *loggers.Data, req playListModel.ReqPlayList) (data *playListModel.PlayList, err error)
	DeletePlayList(record *loggers.Data, id int64) (err error)
}

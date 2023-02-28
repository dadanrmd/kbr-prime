package playListRepository

import (
	"kbrprime-be/internal/app/model/playListModel"
)

type IPlayListRepository interface {
	FindByUser(UserId int64) (data *[]playListModel.ListPlayList, err error)
	FindById(id int64) (data *playListModel.PlayList, err error)
	Insert(data playListModel.PlayList) (*playListModel.PlayList, error)
	Delete(data playListModel.PlayList) error
}

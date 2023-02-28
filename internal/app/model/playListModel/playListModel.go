package playListModel

import (
	"kbrprime-be/internal/app/model/helperModel"
	"kbrprime-be/internal/app/model/podcastModel"
)

/* Table Definition */
type PlayList struct {
	Id                                  int64 `json:"id"`
	UserId                              int64 `json:"user_id"`
	EpisodeId                           int64 `json:"episode_id"`
	helperModel.DateAuditModelTimeStamp `json:"-"`
}

func (PlayList) TableName() string {
	return "kbr_play_list"
}

// DTO

type ReqPlayList struct {
	UserId    int64 `json:"user_id"`
	EpisodeId int64 `json:"episode_id"`
}

type ListPlayList struct {
	IdPlayList int64 `json:"id_play_list"`
	podcastModel.Podcast
}

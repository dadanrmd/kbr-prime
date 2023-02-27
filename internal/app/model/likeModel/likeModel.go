package likeModel

import "kbrprime-be/internal/app/model/helperModel"

/* Table Definition */
type Like struct {
	Id                                  int64 `json:"id"`
	UserId                              int64 `json:"user_id"`
	EpisodeId                           int64 `json:"episode_id"`
	helperModel.DateAuditModelTimeStamp `json:"-"`
}

func (Like) TableName() string {
	return "kbr_like"
}

type ReqLike struct {
	UserId    int64 `json:"user_id"`
	EpisodeId int64 `json:"episode_id"`
}

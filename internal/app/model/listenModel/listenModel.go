package listenModel

import "kbrprime-be/internal/app/model/helperModel"

/* Table Definition */
type Listen struct {
	Id                                  int64  `json:"id"`
	UserId                              int64  `json:"user_id"`
	SessionId                           string `json:"session_id"`
	EpisodeId                           int64  `json:"episode_id"`
	Play                                int64  `json:"play"`
	helperModel.DateAuditModelTimeStamp `json:"-"`
}

func (Listen) TableName() string {
	return "kbr_listen"
}

/* DTO Definition */
type ReqRecord struct {
	UserId    int64  `json:"user_id"`
	SessionId string `json:"session_id"`
	EpisodeId int64  `json:"episode_id"`
	Play      int64  `json:"play"`
}

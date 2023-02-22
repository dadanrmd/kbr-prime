package healtyModel

import "kbrprime-be/internal/app/model/helperModel"

/* Table Definition */
type Healty struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
	helperModel.DateAuditModel
	helperModel.UserAuditModel
}

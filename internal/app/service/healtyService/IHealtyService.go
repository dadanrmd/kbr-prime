package healtyService

import (
	"kbrprime-be/internal/app/model/healtyModel"

	"kbrprime-be/internal/app/commons/loggers"
)

type IHealtyService interface {
	FindAllHealty(record *loggers.Data) (*[]healtyModel.Healty, error)
}

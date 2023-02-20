package healtyService

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/model/healtyModel"

	"github.com/rizkianakbar/kbrprime-be/internal/app/commons/loggers"
)

type IHealtyService interface {
	FindAllHealty(record *loggers.Data) (*[]healtyModel.Healty, error)
}

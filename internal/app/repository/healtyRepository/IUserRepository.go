package healtyRepository

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/model/healtyModel"
)

type IHealtyRepository interface {
	FindAll() (*[]healtyModel.Healty, error)
}

package healtyRepository

import "kbrprime-be/internal/app/model/healtyModel"

type IHealtyRepository interface {
	FindAll() (*[]healtyModel.Healty, error)
}

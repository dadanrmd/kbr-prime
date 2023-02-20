package healtyRepository

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/model/healtyModel"

	"gorm.io/gorm"
)

type healtyRepository struct {
	db *gorm.DB
}

func NewHealtyRepository(db *gorm.DB) IHealtyRepository {
	return &healtyRepository{db: db}
}
func (h healtyRepository) FindAll() (*[]healtyModel.Healty, error) {
	var data []healtyModel.Healty
	db := h.db.Find(&data)
	return &data, db.Error
}

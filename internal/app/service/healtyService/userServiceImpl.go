package healtyService

import (
	"errors"

	"github.com/rizkianakbar/kbrprime-be/internal/app/model/healtyModel"
	"github.com/rizkianakbar/kbrprime-be/internal/app/repository/healtyRepository"

	"github.com/rizkianakbar/kbrprime-be/internal/app/commons/loggers"
)

var (
	ErrHealtyNotFound = errors.New("healty not found")
)

type healtyService struct {
	healtyRepo healtyRepository.IHealtyRepository
}

func NewHealtyService(healty healtyRepository.IHealtyRepository) IHealtyService {
	return &healtyService{healtyRepo: healty}
}

func (h healtyService) FindAllHealty(record *loggers.Data) (*[]healtyModel.Healty, error) {
	loggers.Logf(record, "Info, FindAllHealty")
	data, err := h.healtyRepo.FindAll()
	if err != nil {
		return nil, ErrHealtyNotFound
	}
	return data, err
}

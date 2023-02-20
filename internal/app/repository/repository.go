package repository

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/commons"
	"github.com/rizkianakbar/kbrprime-be/internal/app/repository/healtyRepository"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
}

type Repositories struct {
	HealtyRepository healtyRepository.IHealtyRepository
}

package repository

import (
	"kbrprime-be/internal/app/commons"
	"kbrprime-be/internal/app/repository/healtyRepository"
	"kbrprime-be/internal/app/repository/showRepository"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
}

type Repositories struct {
	HealtyRepository healtyRepository.IHealtyRepository
	ShowRepository   showRepository.IShowRepository
}

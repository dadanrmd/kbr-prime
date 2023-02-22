package service

import (
	"kbrprime-be/internal/app/commons"
	"kbrprime-be/internal/app/repository"
	"kbrprime-be/internal/app/service/healtyService"
	"kbrprime-be/internal/app/service/showService"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repositories
}

type Services struct {
	HealtyService healtyService.IHealtyService
	ShowService   showService.IShowService
}

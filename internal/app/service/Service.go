package service

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/commons"
	"github.com/rizkianakbar/kbrprime-be/internal/app/repository"
	"github.com/rizkianakbar/kbrprime-be/internal/app/service/healtyService"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repositories
}

type Services struct {
	HealtyService healtyService.IHealtyService
}

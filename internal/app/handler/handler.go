package handler

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/commons"
	"github.com/rizkianakbar/kbrprime-be/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}

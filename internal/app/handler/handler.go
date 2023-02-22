package handler

import (
	"kbrprime-be/internal/app/commons"
	"kbrprime-be/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}

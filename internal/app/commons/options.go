package commons

import (
	"github.com/rizkianakbar/kbrprime-be/internal/app/appcontext"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	AppCtx    *appcontext.AppContext
	Db        *gorm.DB
	UUID      Iuuid
	Validator *validator.Validate
}

package categoriesService

import (
	"kbrprime-be/internal/app/model/categoriesModel"

	"kbrprime-be/internal/app/commons/loggers"
)

type ICategoriesService interface {
	FindAllCategories(record *loggers.Data) (data *[]categoriesModel.Categories, err error)
}

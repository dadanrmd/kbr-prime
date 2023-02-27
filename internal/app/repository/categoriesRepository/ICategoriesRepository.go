package categoriesRepository

import (
	"kbrprime-be/internal/app/model/categoriesModel"
)

type ICategoriesRepository interface {
	FindAll() (data *[]categoriesModel.Categories, err error)
	FindByCategoryScope(categoryScope string) (data *[]categoriesModel.Categories, err error)
}

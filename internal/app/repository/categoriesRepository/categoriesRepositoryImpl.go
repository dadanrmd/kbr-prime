package categoriesRepository

import (
	"kbrprime-be/internal/app/model/categoriesModel"

	"gorm.io/gorm"
)

type categoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) ICategoriesRepository {
	return &categoriesRepository{db: db}
}
func (h categoriesRepository) FindAll() (data *[]categoriesModel.Categories, err error) {
	err = h.db.Find(&data).Error
	return
}
func (h categoriesRepository) FindByCategoryScope(categoryScope string) (data *[]categoriesModel.Categories, err error) {
	err = h.db.Where("category_scope = ?", categoryScope).Find(&data).Error
	return
}

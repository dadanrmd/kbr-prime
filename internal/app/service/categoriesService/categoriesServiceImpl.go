package categoriesService

import (
	"errors"
	"kbrprime-be/internal/app/model/categoriesModel"
	"kbrprime-be/internal/app/repository/categoriesRepository"

	"kbrprime-be/internal/app/commons/loggers"
)

var (
	errKategori = errors.New("Katgori tidak ada")
)

type categoriesService struct {
	categoriesRepo categoriesRepository.ICategoriesRepository
}

func NewCategoriesService(categoriesRepo categoriesRepository.ICategoriesRepository) ICategoriesService {
	return &categoriesService{categoriesRepo}
}

func (h categoriesService) FindAllCategories(record *loggers.Data) (data *[]categoriesModel.Categories, err error) {
	loggers.Logf(record, "Info, FindAllCategories")
	// data, err = h.categoriesRepo.FindAll()
	data, err = h.categoriesRepo.FindByCategoryScope("News")
	if err != nil {
		loggers.Logf(record, "Error, FindAll")
		err = errKategori
		return
	}
	return
}

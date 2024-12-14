package usecase

import (
	"github.com/CracherX/catalog_hist/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepository
}

func NewProductUseCase(repo ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

// GetProducts возвращает товары с учетом пагинации
func (uc *ProductUseCase) GetProducts(page, pageSize, from, untill int, countries, categories []string) ([]entity.Product, int64, error) {
	offset := (page - 1) * pageSize
	prod, err := uc.repo.GetProducts(pageSize, offset, from, untill, countries, categories)
	if err != nil {
		return nil, 0, err
	}
	count, err := uc.repo.CountRecords(from, untill, countries, categories)
	if err != nil {
		return nil, 0, err
	}
	return prod, count, nil
}

func (uc *ProductUseCase) GetConcreteProduct(id int) (*entity.Product, error) {
	return uc.repo.GetProduct(id)
}

// UpdateProduct обрабатывает обновление продукта
func (uc *ProductUseCase) UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error) {

	product, err := uc.repo.UpdateProduct(id, updates)
	if err != nil {
		return nil, err
	}

	return product, nil
}

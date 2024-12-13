package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	GetProducts(limit, offset int, countryName, categoryName string) ([]entity.Product, error) // Получение списка товаров
	CountRecords(countryName, categoryName string) (int64, error)
	GetProduct(id int) (*entity.Product, error)
}

package usecase

// ProductRepository — интерфейс для работы с продуктами
type ProductRepository interface {
	GetProducts(limit, offset int) ([]entity.Product, error) // Получение списка товаров
}

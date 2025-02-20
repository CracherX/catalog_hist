package repository

import (
	"github.com/CracherX/catalog_hist/internal/entity"
	"gorm.io/gorm"
)

type ProductRepoGorm struct {
	db *gorm.DB
}

func NewProductRepoGorm(db *gorm.DB) *ProductRepoGorm {
	return &ProductRepoGorm{db: db}
}

// GetProducts возвращает список товаров с пагинацией
func (r *ProductRepoGorm) GetProducts(limit, offset, until, from int, countries, categories []string) ([]entity.Product, error) {
	var products []entity.Product
	query := r.db.Preload("Country").Preload("Category")

	// Фильтрация по странам
	if len(countries) != 1 || countries[0] != "all" {
		query = query.Joins("JOIN countries ON countries.id = products.country_id").
			Where("countries.name IN ?", countries)
	}

	// Фильтрация по категориям
	if len(categories) != 1 || categories[0] != "all" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name IN ?", categories)
	}

	// Фильтрация по диапазону дат
	query = query.Where("products.year BETWEEN ? AND ?", until, from)

	// Пагинация
	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepoGorm) CountRecords(from, untill int, countries, categories []string) (int64, error) {
	var count int64
	var products entity.Product
	query := r.db.Model(&products)

	// Фильтрация по странам
	if len(countries) != 1 || countries[0] != "all" {
		query = query.Joins("JOIN countries ON countries.id = products.country_id").
			Where("countries.name IN ?", countries)
	}

	// Фильтрация по категориям
	if len(categories) != 1 || categories[0] != "all" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name IN ?", categories)
	}

	// Фильтрация по диапазону дат
	query = query.Where("products.year BETWEEN ? AND ?", from, untill)

	// Подсчёт количества
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetProduct возвращает данные о конкретном товаре
func (r *ProductRepoGorm) GetProduct(id int) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Preload("Country").Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepoGorm) UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.Preload("Country").Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}

	// Обновление указанных полей
	if err := r.db.Model(&product).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepoGorm) DeleteProduct(id int) (*entity.Product, error) {
	// Находим продукт по ID
	var product entity.Product
	if err := r.db.Preload("Country").Preload("Category").First(&product, id).Error; err != nil {
		// Если продукт не найден, возвращаем ошибку
		return nil, err
	}

	// Удаляем продукт из базы данных
	if err := r.db.Delete(&product).Error; err != nil {
		// Возвращаем ошибку, если не удалось удалить
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepoGorm) AddProduct(product *entity.Product) (*entity.Product, error) {
	// Добавляем продукт в базу данных
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	// Загружаем связанные данные (Country и Category) для возврата полного объекта
	if err := r.db.Preload("Country").Preload("Category").First(product, product.ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}

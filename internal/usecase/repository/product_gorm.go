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
func (r *ProductRepoGorm) GetProducts(limit, offset int, countryName, categoryName string) ([]entity.Product, error) {
	var products []entity.Product
	query := r.db.Preload("Country").Preload("Category")

	// Добавляем фильтрацию по стране, если указана
	if countryName != "all" {
		query = query.Where("countries.name = ?", countryName).Joins("JOIN countries ON countries.id = products.country_id")
	}

	// Добавляем фильтрацию по категории, если указана
	if categoryName != "all" {
		query = query.Where("categories.name = ?", categoryName).Joins("JOIN categories ON categories.id = products.category_id")
	}

	// Применяем пагинацию
	query = query.Limit(limit).Offset(offset)

	// Выполняем запрос
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepoGorm) CountRecords(countryName, categoryName string) (int64, error) {
	var products entity.Product
	var count int64
	query := r.db.Model(&products)
	if countryName != "all" {
		query = query.Where("countries.name = ?", countryName).Joins("JOIN countries ON countries.id = products.country_id")
	}
	if categoryName != "all" {
		query = query.Where("categories.name = ?", categoryName).Joins("JOIN categories ON categories.id = products.category_id")
	}
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

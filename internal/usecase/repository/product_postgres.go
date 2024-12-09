package repository

import (
	"catalog/internal/entity"
	"gorm.io/gorm"
)

type ProductRepoGorm struct {
	db *gorm.DB
}

func NewProductRepoGorm(db *gorm.DB) *ProductRepoGorm {
	return &ProductRepoGorm{db: db}
}

// GetProducts возвращает список товаров с пагинацией
func (r *ProductRepoGorm) GetProducts(limit, offset int) ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

package repository

import (
	"github.com/CracherX/catalog_hist/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepoGorm struct {
	db *gorm.DB
}

func NewCategoryRepoGorm(db *gorm.DB) *CategoryRepoGorm {
	return &CategoryRepoGorm{db: db}
}

// GetAllCategories возвращает список всех категорий
func (r *CategoryRepoGorm) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

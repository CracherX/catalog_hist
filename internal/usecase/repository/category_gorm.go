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

func (r *CategoryRepoGorm) UpdateCategory(id int, updates map[string]interface{}) (*entity.Category, error) {
	var product entity.Category
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	// Обновление указанных полей
	if err := r.db.Model(&product).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *CategoryRepoGorm) DeleteCategory(id int) error {
	// Находим продукт по ID
	var product entity.Category
	if err := r.db.First(&product, id).Error; err != nil {
		// Если продукт не найден, возвращаем ошибку
		return err
	}

	// Удаляем продукт из базы данных
	if err := r.db.Delete(&product).Error; err != nil {
		// Возвращаем ошибку, если не удалось удалить
		return err
	}

	return nil
}

func (r *CategoryRepoGorm) AddCategory(product *entity.Category) (*entity.Category, error) {
	// Добавляем продукт в базу данных
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	// Загружаем связанные данные (Country и Category) для возврата полного объекта
	if err := r.db.First(product, product.ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}

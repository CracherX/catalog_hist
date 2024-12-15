package repository

import (
	"github.com/CracherX/catalog_hist/internal/entity"
	"gorm.io/gorm"
)

type CountryRepoGorm struct {
	db *gorm.DB
}

func NewCountryRepoGorm(db *gorm.DB) *CountryRepoGorm {
	return &CountryRepoGorm{db: db}
}

// GetAllCountries возвращает список всех стан
func (r *CountryRepoGorm) GetAllCountries() ([]entity.Country, error) {
	var countries []entity.Country
	err := r.db.Find(&countries).Error
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (r *CountryRepoGorm) UpdateCountry(id int, updates map[string]interface{}) (*entity.Country, error) {
	var product entity.Country
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	// Обновление указанных полей
	if err := r.db.Model(&product).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *CountryRepoGorm) DeleteCountry(id int) error {
	// Находим продукт по ID
	var product entity.Country
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

func (r *CountryRepoGorm) AddCountry(product *entity.Country) (*entity.Country, error) {
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

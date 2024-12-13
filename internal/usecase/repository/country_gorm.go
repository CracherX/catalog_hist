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

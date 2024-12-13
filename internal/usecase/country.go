package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

type CountryUseCase struct {
	repo CountryRepository
}

func NewCountryUseCase(repo CountryRepository) *CountryUseCase {
	return &CountryUseCase{repo: repo}
}

func (uc *CountryUseCase) GetCountries() ([]entity.Country, error) {
	return uc.repo.GetAllCountries()
}

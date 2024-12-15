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

// UpdateCategory обрабатывает обновление продукта
func (uc *CountryUseCase) UpdateCategory(id int, updates map[string]interface{}) (*entity.Country, error) {

	product, err := uc.repo.UpdateCountry(id, updates)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *CountryUseCase) DeleteCountry(id int) error {
	return uc.repo.DeleteCountry(id)
}

func (uc *CountryUseCase) AddCountry(product *entity.Country) (*entity.Country, error) {
	// Проверяем валидность данных, если требуется (можно реализовать бизнес-валидацию)

	// Вызываем метод репозитория для сохранения продукта
	newProduct, err := uc.repo.AddCountry(product)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

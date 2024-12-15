package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

type CategoryUseCase struct {
	repo CategoryRepository
}

func NewCategoryUseCase(repo CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{repo: repo}
}

func (uc *CategoryUseCase) GetCategories() ([]entity.Category, error) {
	return uc.repo.GetAllCategories()
}

// UpdateCategory обрабатывает обновление продукта
func (uc *CategoryUseCase) UpdateCategory(id int, updates map[string]interface{}) (*entity.Category, error) {

	product, err := uc.repo.UpdateCategory(id, updates)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *CategoryUseCase) DeleteCategory(id int) error {
	return uc.repo.DeleteCategory(id)
}

func (uc *CategoryUseCase) AddCategory(product *entity.Category) (*entity.Category, error) {
	// Проверяем валидность данных, если требуется (можно реализовать бизнес-валидацию)

	// Вызываем метод репозитория для сохранения продукта
	newProduct, err := uc.repo.AddCategory(product)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

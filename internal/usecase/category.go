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

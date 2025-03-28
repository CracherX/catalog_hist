package usecase

import "github.com/CracherX/catalog_hist/internal/entity"

type PictureUseCase struct {
	repo PictureRepository
}

func NewPictureUC(repo PictureRepository) *PictureUseCase {
	return &PictureUseCase{repo: repo}
}

func (uc *PictureUseCase) AddPictures(prodID int, url ...string) error {
	return uc.repo.AddPictures(prodID, url...)
}

func (uc *PictureUseCase) DeletePicture(id int) error {
	return uc.repo.DeletePicture(id)
}

func (uc *PictureUseCase) GetPictures(prodId int) ([]entity.Picture, error) {
	return uc.repo.GetAllPictures(prodId)
}

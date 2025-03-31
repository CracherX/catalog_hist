package repository

import (
	"github.com/CracherX/catalog_hist/internal/entity"
	"gorm.io/gorm"
)

type PictureRepoGorm struct {
	db *gorm.DB
}

func NewPictureRepoGorm(db *gorm.DB) *PictureRepoGorm {
	return &PictureRepoGorm{db: db}
}

// GetAllPictures возвращает все картинки
func (r *PictureRepoGorm) GetAllPictures(prodId int) ([]entity.Picture, error) {
	var pictures []entity.Picture
	err := r.db.Where("product_id = ?", prodId).Find(&pictures).Error
	if err != nil {
		return nil, err
	}
	return pictures, nil
}

// DeletePicture удаляет картинку по ID
func (r *PictureRepoGorm) DeletePicture(id int) error {
	var picture entity.Picture
	if err := r.db.First(&picture, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&picture).Error; err != nil {
		return err
	}
	return nil
}

// AddPictures добавить список картинок
func (r *PictureRepoGorm) AddPictures(prodID int, url ...string) error {
	var pictures []entity.Picture

	for _, a := range url {
		picture := entity.Picture{
			PictureURL: a,
			ProductID:  prodID,
		}
		pictures = append(pictures, picture)
	}

	err := r.db.Create(pictures).Error
	if err != nil {
		return err
	}

	return nil
}

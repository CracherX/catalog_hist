package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CracherX/catalog_hist/internal/entity"
	"github.com/elastic/go-elasticsearch/v8"
)

type ProductUseCase struct {
	repo ProductRepository
	es   *elasticsearch.Client
}

func NewProductUseCase(repo ProductRepository, es *elasticsearch.Client) *ProductUseCase {
	return &ProductUseCase{repo: repo, es: es}
}

// GetProducts возвращает товары с учетом пагинации
func (uc *ProductUseCase) GetProducts(page, pageSize, from, untill int, countries, categories []string) ([]entity.Product, int64, error) {
	offset := (page - 1) * pageSize
	prod, err := uc.repo.GetProducts(pageSize, offset, from, untill, countries, categories)
	if err != nil {
		return nil, 0, err
	}
	count, err := uc.repo.CountRecords(from, untill, countries, categories)
	if err != nil {
		return nil, 0, err
	}
	return prod, count, nil
}

func (uc *ProductUseCase) GetConcreteProduct(id int) (*entity.Product, error) {
	return uc.repo.GetProduct(id)
}

// UpdateProduct обрабатывает обновление продукта
func (uc *ProductUseCase) UpdateProduct(id int, updates map[string]interface{}) (*entity.Product, error) {

	product, err := uc.repo.UpdateProduct(id, updates)
	if err != nil {
		return nil, err
	}

	updateIndex(product, uc.es)

	return product, nil
}

func (uc *ProductUseCase) DeleteProduct(id int) error {
	product, err := uc.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	deleteIndex(product, uc.es)

	return nil
}

func (uc *ProductUseCase) AddProduct(product *entity.Product) (*entity.Product, error) {
	newProduct, err := uc.repo.AddProduct(product)
	if err != nil {
		return nil, err
	}

	updateIndex(newProduct, uc.es)

	return newProduct, nil
}

/*
TODO: из за своей тупости прокинуть сюда логер не так просто будет,

	а точнее просто, но процесс логирования ты решил оставить на handlers,
	короче как всегда вьебал архитектуру, можешь перекинуть процесс индексации
	в handlers в целом, но там чето тогда слишком много говна будет, оставь всё так
*/
func updateIndex(product *entity.Product, es *elasticsearch.Client) {
	go func(product *entity.Product, es *elasticsearch.Client) {
		data, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err.Error())
		}
		docID := fmt.Sprintf("%d", product.ID)
		_, err = es.Index("products", bytes.NewReader(data), es.Index.WithDocumentID(docID))
		if err != nil {
			fmt.Println(err.Error())
		}
	}(product, es)
}

func deleteIndex(product *entity.Product, es *elasticsearch.Client) {
	go func(product *entity.Product, es *elasticsearch.Client) {
		docID := fmt.Sprintf("%d", product.ID)
		_, err := es.Delete("products", docID)
		if err != nil {
			fmt.Println(err.Error())
		}
	}(product, es)
}

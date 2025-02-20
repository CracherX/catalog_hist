package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CracherX/catalog_hist/internal/controller/http/handlers"
	"github.com/CracherX/catalog_hist/internal/entity"
	"github.com/CracherX/catalog_hist/pkg/config"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config, db *gorm.DB, log handlers.Logger) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(parseConfig(cfg.ElasticClient.Url))
	if err != nil {
		return nil, err
	}
	go indexProducts(client, db, log)
	return client, nil
}

func parseConfig(url string) elasticsearch.Config {
	return elasticsearch.Config{Addresses: []string{url}}
}

func indexProducts(es *elasticsearch.Client, db *gorm.DB, log handlers.Logger) {
	var products []entity.Product
	if err := db.Preload("Country").Preload("Category").Find(&products).Error; err != nil {
		log.Error("Ошибка получения товаров из БД", "Ошибка", err.Error())
	}

	for _, product := range products {
		data, _ := json.Marshal(product)

		docID := fmt.Sprintf("%d", product.ID)

		// Индексируем или обновляем документ в Elasticsearch
		_, err := es.Index("products", bytes.NewReader(data), es.Index.WithDocumentID(docID))
		if err != nil {
			log.Error("Ошибка индексации в поисковой движок", "ID записи", product.ID, "Ошибка", err.Error())
		}
	}
}

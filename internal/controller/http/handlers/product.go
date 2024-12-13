package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	uc  ProductUseCase
	val Validator
	log Logger
}

func NewProductHandler(uc ProductUseCase, val Validator, log Logger) *ProductHandler {
	return &ProductHandler{
		uc:  uc,
		val: val,
		log: log,
	}
}

func (ph *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	data := dto.GetProductsRequest{
		PageSize: vars.Get("pageSize"),
		Page:     vars.Get("page"),
		Category: vars.Get("category"),
		Country:  vars.Get("country"),
	}
	if data.PageSize == "" {
		data.PageSize = "15"
	}
	if data.Page == "" {
		data.Page = "1"
	}
	if data.Category == "" {
		data.Category = "all"
	}
	if data.Country == "" {
		data.Country = "all"
	}
	err := ph.val.Validate(&data)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	pageSize, err := strconv.Atoi(data.PageSize)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	page, err := strconv.Atoi(data.Page)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	prod, count, err := ph.uc.GetProducts(page, pageSize, data.Country, data.Category)
	if err != nil {
		ph.log.Error("Ошибка работы UseCase", "Запрос", "GetProducts", "Ошибка", err.Error())
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническом специалисту")
		}
		return
	}

	prds := dto.ToProductDTOs(prod)
	resp := dto.ToProductsDTO(prds, count, page, pageSize, data.Category, data.Country)

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		ph.log.Error("Ошибка работы энкодера", "Запрос", "GetProducts", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническом специалисту")
		return
	}
}

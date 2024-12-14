package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"net/http"
	"strconv"
	"time"
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
		From:     vars.Get("from"),
		Until:    vars.Get("until"),
	}

	_ = json.NewDecoder(r.Body).Decode(&data)

	// Установка значений по умолчанию
	if data.PageSize == "" {
		data.PageSize = "15"
	}
	if data.Page == "" {
		data.Page = "1"
	}
	if data.From == "" {
		data.From = "1900" // Минимальная дата по умолчанию
	}
	if data.Until == "" {
		data.Until = strconv.Itoa(time.Now().Year()) // Текущий год по умолчанию
	}
	if len(data.Categories) == 0 {
		data.Categories = []string{"all"}
	}
	if len(data.Countries) == 0 {
		data.Countries = []string{"all"}
	}

	// Валидация DTO
	err := ph.val.Validate(&data)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	// Преобразование параметров
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
	from, err := strconv.Atoi(data.From)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	untill, err := strconv.Atoi(data.Until)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProducts")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	// Вызов UseCase
	prod, count, err := ph.uc.GetProducts(page, pageSize, from, untill, data.Countries, data.Categories)
	if err != nil {
		ph.log.Error("Ошибка работы UseCase", "Запрос", "GetProducts", "Ошибка", err.Error())
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		}
		return
	}

	prds := dto.ToProductDTOs(prod)
	resp := dto.ToProductsDTO(prds, count, page, pageSize, data.Categories, data.Countries)

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		ph.log.Error("Ошибка работы энкодера", "Запрос", "GetProducts", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	data := dto.GetProductRequest{
		ID: vars.Get("id"),
	}
	err := ph.val.Validate(&data)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProductById")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	id, err := strconv.Atoi(data.ID)
	if err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "GetProductById")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}
	prod, err := ph.uc.GetConcreteProduct(id)
	if err != nil {
		ph.log.Error("Ошибка работы UseCase", "Запрос", "GetProducts", "Ошибка", err.Error())
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		}
		return
	}
	res := dto.ToProductDTO(prod)

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		ph.log.Error("Ошибка работы энкодера", "Запрос", "GetProducts", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

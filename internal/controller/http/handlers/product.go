package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"github.com/CracherX/catalog_hist/internal/entity"
	"net/http"
	"strconv"
	"time"
)

type ProductHandler struct {
	uc  ProductUseCase
	val Validator
	log Logger
	cl  Client
}

func NewProductHandler(uc ProductUseCase, val Validator, log Logger, cl Client) *ProductHandler {
	return &ProductHandler{
		uc:  uc,
		val: val,
		log: log,
		cl:  cl,
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
	err = json.NewEncoder(w).Encode(&resp)
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

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		ph.log.Error("Ошибка работы энкодера", "Запрос", "GetProducts", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

func (ph *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "PatchProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data := dto.PatchProductRequest{
		JWT:     query.Get("jwt"),
		ID:      query.Get("id"),
		Updates: updates,
	}
	if err := ph.val.Validate(&data); err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "PatchProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ph.auth(data.JWT, w); err != nil {
		return
	}

	id, _ := strconv.Atoi(data.ID)

	_, err := ph.uc.UpdateProduct(id, data.Updates)
	if err != nil {
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно (еблан, указывай id страны и категории)")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Обновление выполнено успешно!")
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data := dto.DeleteProductRequest{JWT: query.Get("jwt"), ID: query.Get("id")}
	if err := ph.val.Validate(&data); err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "DeleteProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ph.auth(data.JWT, w); err != nil {
		return
	}

	id, _ := strconv.Atoi(data.ID)
	if err := ph.uc.DeleteProduct(id); err != nil {
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Удаление выполнено успешно!")
}

func (ph *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var data dto.AddProductRequest
	query := r.URL.Query()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "AddProduct", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data.JWT = query.Get("jwt")

	if err := ph.val.Validate(&data); err != nil {
		ph.log.Debug("Получен Bad Request", "Запрос", "AddProduct", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ph.auth(data.JWT, w); err != nil {
		return
	}

	product := entity.Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Year:        data.Year,
		CountryID:   data.CountryId,
		CategoryID:  data.CategoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := ph.uc.AddProduct(&product)
	if err != nil {
		if errors.Is(err, driver.ErrBadConn) {
			ph.log.Error("Ошибка работы базы данных", "Запрос", "GetProduct", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			ph.log.Debug("Повторение уникального ключа", "Запрос", "GetProduct", "Ошибка", err.Error())
			dto.Response(w, http.StatusConflict, "Conflict", "Повторение уникальных значений")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusCreated, "Добавление выполнено успешно!")
}

func (ph *ProductHandler) auth(jwt string, w http.ResponseWriter) error {
	var cdto dto.AuthClientResponse
	params := map[string]string{
		"jwt": jwt,
	}

	clr, err := ph.cl.Get("/auth/profile", params)
	if err != nil {
		ph.log.Error("Ошибка в работе клиента", "Запрос", "auth")
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе внешних сервисов")
		return err
	}

	if err = json.NewDecoder(clr.Body).Decode(&cdto); err != nil {
		ph.log.Error("Ошибка работы энкодера", "Запрос", "auth", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return err
	}

	if cdto.IsAdmin != true {
		ph.log.Debug("Недостаточно прав для выполнения запроса", "Запрос", "auth")
		dto.Response(w, http.StatusForbidden, "Forbidden", "У вас недостаточно прав для вызова данного метода")
		return errors.New("недостаточно прав пользователя")
	}
	return nil
}

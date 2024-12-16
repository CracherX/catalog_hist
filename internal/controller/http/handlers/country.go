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

type CountryHandler struct {
	uc  CountryUseCase
	val Validator
	log Logger
	cl  Client
}

func NewCountryHandler(uc CountryUseCase, val Validator, log Logger, cl Client) *CountryHandler {
	return &CountryHandler{
		uc:  uc,
		val: val,
		log: log,
		cl:  cl,
	}
}

func (ch *CountryHandler) GetCountries(w http.ResponseWriter, _ *http.Request) {
	count, err := ch.uc.GetCountries()
	if err != nil {
		ch.log.Error("Ошибка работы UseCase", "Запрос", "GetCountries", "Ошибка", err.Error())
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		}
		return
	}
	res := dto.ToCountriesDTO(count)

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)

	if err != nil {
		ch.log.Error("Ошибка работы энкодера", "Запрос", "GetCategories", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

func (ch *CountryHandler) PatchCountry(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "PatchProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data := dto.PatchRequest{
		JWT:     query.Get("jwt"),
		ID:      query.Get("id"),
		Updates: updates,
	}
	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "PatchProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ch.auth(data.JWT, w); err != nil {
		return
	}

	id, _ := strconv.Atoi(data.ID)

	_, err := ch.uc.UpdateCategory(id, data.Updates)
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

func (ch *CountryHandler) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data := dto.DeleteRequest{JWT: query.Get("jwt"), ID: query.Get("id")}
	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "DeleteProduct")
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ch.auth(data.JWT, w); err != nil {
		return
	}

	id, _ := strconv.Atoi(data.ID)
	if err := ch.uc.DeleteCountry(id); err != nil {
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Удаление выполнено успешно!")
}

func (ch *CountryHandler) AddCountry(w http.ResponseWriter, r *http.Request) {
	var data dto.AddCountryRequest
	query := r.URL.Query()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "AddProduct", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data.JWT = query.Get("jwt")

	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "AddProduct", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ch.auth(data.JWT, w); err != nil {
		return
	}

	product := entity.Country{
		Name:      data.Name,
		Code:      data.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := ch.uc.AddCountry(&product)
	if err != nil {
		if errors.Is(err, driver.ErrBadConn) {
			ch.log.Error("Ошибка работы базы данных", "Запрос", "GetProduct", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			ch.log.Debug("Повторение уникального ключа", "Запрос", "GetProduct", "Ошибка", err.Error())
			dto.Response(w, http.StatusConflict, "Conflict", "Повторение уникальных значений")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	dto.Response(w, http.StatusCreated, "Добавление выполнено успешно!")
}

func (ch *CountryHandler) auth(jwt string, w http.ResponseWriter) error {
	var cdto dto.AuthClientResponse
	params := map[string]string{
		"jwt": jwt,
	}

	clr, err := ch.cl.Get("/auth/profile", params)
	if err != nil {
		ch.log.Error("Ошибка в работе клиента", "Запрос", "auth")
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе внешних сервисов")
		return err
	}

	if err = json.NewDecoder(clr.Body).Decode(&cdto); err != nil {
		ch.log.Error("Ошибка работы энкодера", "Запрос", "auth", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return err
	}

	if cdto.IsAdmin != true {
		ch.log.Debug("Недостаточно прав для выполнения запроса", "Запрос", "auth")
		dto.Response(w, http.StatusForbidden, "Forbidden", "У вас недостаточно прав для вызова данного метода")
		return errors.New("недостаточно прав пользователя")
	}
	return nil
}

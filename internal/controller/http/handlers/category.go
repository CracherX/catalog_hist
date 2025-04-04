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

type CategoryHandler struct {
	uc  CategoryUseCase
	val Validator
	log Logger
	cl  Client
}

func NewCategoryHandler(uc CategoryUseCase, val Validator, log Logger, cl Client) *CategoryHandler {
	return &CategoryHandler{
		uc:  uc,
		val: val,
		log: log,
		cl:  cl,
	}
}

func (ch *CategoryHandler) GetCategories(w http.ResponseWriter, _ *http.Request) {
	cat, err := ch.uc.GetCategories()
	if err != nil {
		ch.log.Error("Ошибка работы UseCase", "Запрос", "GetProducts", "Ошибка", err.Error())
		if errors.Is(err, driver.ErrBadConn) {
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		}
		return
	}
	res := dto.ToCategoriesDTO(cat)

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)

	if err != nil {
		ch.log.Error("Ошибка работы энкодера", "Запрос", "GetCategories", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

func (ch *CategoryHandler) PatchCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "PatchCategory", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data := dto.PatchRequest{
		JWT:     query.Get("jwt"),
		ID:      query.Get("id"),
		Updates: updates,
	}
	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "PatchCategory", "Ошибка", err.Error())
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
			ch.log.Error("BadGateway", "Запрос", "PatchCategory", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			ch.log.Error("BadGateway", "Запрос", "PatchCategory", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно (еблан, указывай id страны и категории)")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Обновление выполнено успешно!")
}

func (ch *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data := dto.DeleteRequest{
		JWT: query.Get("jwt"),
		ID:  query.Get("id"),
	}
	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "DeleteCategory", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ch.auth(data.JWT, w); err != nil {
		return
	}

	id, _ := strconv.Atoi(data.ID)
	if err := ch.uc.DeleteCategory(id); err != nil {
		ch.log.Error("Bad Gateway", "Запрос", "DeleteCategory", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Удаление выполнено успешно!")
}

func (ch *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var data dto.AddCategoryRequest
	query := r.URL.Query()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "AddCategory", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	data.JWT = query.Get("jwt")

	if err := ch.val.Validate(&data); err != nil {
		ch.log.Debug("Получен Bad Request", "Запрос", "AddCategory", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err := ch.auth(data.JWT, w); err != nil {
		return
	}

	product := entity.Category{
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := ch.uc.AddCategory(&product)
	if err != nil {
		if errors.Is(err, driver.ErrBadConn) {
			ch.log.Error("Ошибка работы базы данных", "Запрос", "AddCategory", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Ошибка в работе внешних сервисов")
		} else {
			ch.log.Debug("Повторение уникального ключа", "Запрос", "AddCategory", "Ошибка", err.Error())
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

func (ch *CategoryHandler) auth(jwt string, w http.ResponseWriter) error {
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

package handlers

import (
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"net/http"
	"strconv"
)

type PictureHandler struct {
	uc  PictureUseCase
	val Validator
	log Logger
	cl  Client
}

func NewPictureHandler(uc PictureUseCase, val Validator, log Logger, cl Client) *PictureHandler {
	return &PictureHandler{
		uc:  uc,
		val: val,
		log: log,
		cl:  cl,
	}
}

func (h *PictureHandler) GetAllProductPictures(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data := dto.GetPicturesRequest{
		ProductID: query.Get("productID"),
	}

	err := h.val.Validate(&data)

	if err != nil {
		h.log.Debug("Получен Bad Request", "Запрос", "GetAllProductPictures", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	prid, _ := strconv.Atoi(data.ProductID)

	res, err := h.uc.GetPictures(prid)
	if err != nil {
		h.log.Error("Ошибка", "Запрос", "GetAllProductPictures", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Ошибка какая то я не ебу", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	resdat := dto.GetPicturesResponse{
		ProductID: prid,
		Pictures:  res,
	}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&resdat)
	if err != nil {
		h.log.Error("Ошибка работы энкодера", "Запрос", "GetProducts", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

func (h *PictureHandler) DeletePicture(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	data := dto.DeleteRequest{
		JWT: query.Get("jwt"),
		ID:  query.Get("ID"),
	}

	err := h.val.Validate(&data)
	if err != nil {
		h.log.Debug("Получен Bad Request", "Запрос", "DeletePicture", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	id, _ := strconv.Atoi(data.ID)

	if err = h.auth(data.JWT, w); err != nil {
		return
	}

	err = h.uc.DeletePicture(id)
	if err != nil {
		h.log.Debug("Ошибка", "Запрос", "DeletePicture", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Успешное удаление")
}

func (h *PictureHandler) AddPictures(w http.ResponseWriter, r *http.Request) {
	var data dto.AddPicturesRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.log.Debug("Ошибка", "Запрос", "AddPictures", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	err = h.val.Validate(&data)
	if err != nil {
		h.log.Debug("Ошибка", "Запрос", "AddPictures", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	if err = h.auth(data.JWT, w); err != nil {
		return
	}

	err = h.uc.AddPictures(data.ProductID, data.PictureURL...)
	if err != nil {
		h.log.Debug("Ошибка", "Запрос", "AddPictures", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Ошибка какая то я не ебу тоже", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	dto.Response(w, http.StatusCreated, "Успех!", "Фотографии добавлены!")
}

func (h *PictureHandler) auth(jwt string, w http.ResponseWriter) error {
	var cdto dto.AuthClientResponse
	params := map[string]string{
		"jwt": jwt,
	}

	clr, err := h.cl.Get("/auth/profile", params)
	if err != nil {
		h.log.Error("Ошибка в работе клиента", "Запрос", "auth")
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе внешних сервисов")
		return err
	}

	if err = json.NewDecoder(clr.Body).Decode(&cdto); err != nil {
		h.log.Error("Ошибка работы энкодера", "Запрос", "auth", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return err
	}

	if cdto.IsAdmin != true {
		h.log.Debug("Недостаточно прав для выполнения запроса", "Запрос", "auth")
		dto.Response(w, http.StatusForbidden, "Forbidden", "У вас недостаточно прав для вызова данного метода")
		return errors.New("недостаточно прав пользователя")
	}
	return nil
}

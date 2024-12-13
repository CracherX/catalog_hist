package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"net/http"
)

type CategoryHandler struct {
	uc  CategoryUseCase
	val Validator
	log Logger
}

func NewCategoryHandler(uc CategoryUseCase, val Validator, log Logger) *CategoryHandler {
	return &CategoryHandler{
		uc:  uc,
		val: val,
		log: log,
	}
}

func (ch *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
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

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		ch.log.Error("Ошибка работы энкодера", "Запрос", "GetCategories", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return
	}
}

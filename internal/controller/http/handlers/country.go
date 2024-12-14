package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/catalog_hist/internal/controller/http/dto"
	"net/http"
)

type CountryHandler struct {
	uc  CountryUseCase
	val Validator
	log Logger
}

func NewCountryHandler(uc CountryUseCase, val Validator, log Logger) *CountryHandler {
	return &CountryHandler{
		uc:  uc,
		val: val,
		log: log,
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

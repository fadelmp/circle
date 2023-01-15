package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type LocationController struct {
	LocationUsecase usecase.LocationUsecase
}

func ProviderLocationController(l usecase.LocationUsecase) LocationController {
	return LocationController{LocationUsecase: l}
}

func (l *LocationController) GetAllCountry(e echo.Context) error {

	res := l.LocationUsecase.GetAllCountry()

	return CheckResponse(e, res)
}

func (l *LocationController) GetCountryByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationUsecase.GetCountryByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetAllProvince(e echo.Context) error {

	res := l.LocationUsecase.GetAllProvince()

	return CheckResponse(e, res)
}

func (l *LocationController) GetProvinceByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationUsecase.GetProvinceByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetProvinceByCountryID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("country_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationUsecase.GetProvinceByCountryID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetAllCity(e echo.Context) error {

	res := l.LocationUsecase.GetAllCity()

	return CheckResponse(e, res)
}

func (l *LocationController) GetCityByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationUsecase.GetCityByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetCityByProvinceID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("province_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationUsecase.GetCityByProvinceID(uint(id))

	return CheckResponse(e, res)
}

package controller

import (
	"core/config"
	"core/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type LocationController struct {
	LocationService service.LocationService
}

func ProviderLocationController(l service.LocationService) LocationController {
	return LocationController{LocationService: l}
}

func (l *LocationController) GetAllCountry(e echo.Context) error {

	res := l.LocationService.GetAllCountry()

	return CheckResponse(e, res)
}

func (l *LocationController) GetCountryByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationService.GetCountryByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetAllProvince(e echo.Context) error {

	res := l.LocationService.GetAllProvince()

	return CheckResponse(e, res)
}

func (l *LocationController) GetProvinceByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationService.GetProvinceByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetProvinceByCountryID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("country_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationService.GetProvinceByCountryID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetAllCity(e echo.Context) error {

	res := l.LocationService.GetAllCity()

	return CheckResponse(e, res)
}

func (l *LocationController) GetCityByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationService.GetCityByID(uint(id))

	return CheckResponse(e, res)
}

func (l *LocationController) GetCityByProvinceID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("province_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := l.LocationService.GetCityByProvinceID(uint(id))

	return CheckResponse(e, res)
}

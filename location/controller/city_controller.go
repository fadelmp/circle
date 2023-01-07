package controller

import (
	"location/config"
	"location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type CityController struct {
	CityService service.CityService
}

func ProviderCityController(c service.CityService) CityController {
	return CityController{CityService: c}
}

func (c *CityController) GetAll(e echo.Context) error {

	cities := c.CityService.GetAll()

	if len(cities) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CityNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, cities)
}

func (c *CityController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	city := c.CityService.GetByID(uint(id))
	if city.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CityNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, city)
}

func (c *CityController) GetByprovinceID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("province_id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	cities := c.CityService.GetByprovinceID(uint(id))
	if len(cities) > 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CityNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, cities)
}

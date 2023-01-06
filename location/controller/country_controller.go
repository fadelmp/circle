package controller

import (
	"location/config"
	"location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type CountryController struct {
	CountryService service.CountryService
}

func ProviderCountryController(c service.CountryService) CountryController {
	return CountryController{CountryService: c}
}

func (c *CountryController) GetAll(e echo.Context) error {

	countries := c.CountryService.GetAll()

	if len(countries) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CountryNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, countries)
}

func (c *CountryController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	country := c.CountryService.GetByID(uint(id))
	if country.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CountryNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, country)
}

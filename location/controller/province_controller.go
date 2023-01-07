package controller

import (
	"location/config"
	"location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type ProvinceController struct {
	ProvinceService service.ProvinceService
}

func ProviderProvinceController(p service.ProvinceService) ProvinceController {
	return ProvinceController{ProvinceService: p}
}

func (p *ProvinceController) GetAll(e echo.Context) error {

	provinces := p.ProvinceService.GetAll()

	if len(provinces) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.ProvinceNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, provinces)
}

func (p *ProvinceController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	province := p.ProvinceService.GetByID(uint(id))
	if province.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.ProvinceNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, province)
}

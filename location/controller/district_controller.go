package controller

import (
	"location/config"
	"location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type DistrictController struct {
	DistrictService service.DistrictService
}

func ProviderDistrictController(c service.DistrictService) DistrictController {
	return DistrictController{DistrictService: c}
}

func (d *DistrictController) GetAll(e echo.Context) error {

	sub_districts := d.DistrictService.GetAll()

	if len(sub_districts) == 0 {
		return config.SuccessResponse(e, nil, config.DistrictNotFound)
	}

	return config.SuccessResponse(e, sub_districts, config.GetDistrictSuccess)
}

func (d *DistrictController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	district := d.DistrictService.GetByID(uint(id))

	if district.ID == 0 {
		return config.SuccessResponse(e, nil, config.DistrictNotFound)
	}

	return config.SuccessResponse(e, district, config.GetDistrictSuccess)
}

func (d *DistrictController) GetByCityID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("city_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	districts := d.DistrictService.GetByCityID(uint(id))

	if len(districts) == 0 {
		return config.SuccessResponse(e, nil, config.DistrictNotFound)
	}

	return config.SuccessResponse(e, districts, config.GetDistrictSuccess)
}

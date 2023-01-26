package controller

import (
	"location/config"
	"location/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type SubDistrictController struct {
	SubDistrictService service.SubDistrictService
}

func ProviderSubDistrictController(c service.SubDistrictService) SubDistrictController {
	return SubDistrictController{SubDistrictService: c}
}

func (sd *SubDistrictController) GetAll(e echo.Context) error {

	sub_districts := sd.SubDistrictService.GetAll()

	if len(sub_districts) == 0 {
		return config.SuccessResponse(e, nil, config.SubDistrictNotFound)
	}

	return config.SuccessResponse(e, sub_districts, config.GetSubDistrictSuccess)
}

func (sd *SubDistrictController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	sub_district := sd.SubDistrictService.GetByID(uint(id))

	if sub_district.ID == 0 {
		return config.SuccessResponse(e, nil, config.SubDistrictNotFound)
	}

	return config.SuccessResponse(e, sub_district, config.GetSubDistrictSuccess)
}

func (sd *SubDistrictController) GetByDistrictID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("district_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	sub_districts := sd.SubDistrictService.GetByDistrictID(uint(id))

	if len(sub_districts) == 0 {
		return config.SuccessResponse(e, nil, config.SubDistrictNotFound)
	}

	return config.SuccessResponse(e, sub_districts, config.GetSubDistrictSuccess)
}

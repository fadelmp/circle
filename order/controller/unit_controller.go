package controller

import (
	"net/http"
	"order/config"
	"order/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UnitController struct {
	UnitUsecase usecase.UnitUsecase
}

func ProviderUnitController(s usecase.UnitUsecase) UnitController {
	return UnitController{
		UnitUsecase: s,
	}
}

func (u *UnitController) GetAll(e echo.Context) error {

	units := u.UnitUsecase.GetAll()

	if len(units) == 0 {
		return SuccessResponse(e, nil, config.UnitNotFound)
	}

	return SuccessResponse(e, units, config.GetUnitSuccess)
}

func (u *UnitController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	unit := u.UnitUsecase.GetByID(uint(id))

	if unit.ID == 0 {
		return SuccessResponse(e, nil, config.UnitNotFound)
	}

	return SuccessResponse(e, unit, config.GetUnitSuccess)
}

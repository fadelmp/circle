package controller

import (
	"net/http"
	"order/config"
	"order/dto"
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

func (u *UnitController) Create(e echo.Context) error {

	var unit dto.Unit

	if e.Bind(&unit) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err, err_code := u.UnitUsecase.Create(unit)

	return CheckResponse(e, err, err_code, config.CreateUnitSuccess)
}

func (u *UnitController) Update(e echo.Context) error {

	var unit dto.Unit

	if e.Bind(&unit) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err, err_code := u.UnitUsecase.Update(unit)

	return CheckResponse(e, err, err_code, config.UpdateUnitSuccess)
}

func (u *UnitController) Delete(e echo.Context) error {

	var unit dto.Unit

	if e.Bind(&unit) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := u.UnitUsecase.Delete(unit)

	return CheckResponse(e, err, err_code, config.DeleteUnitSuccess)
}

func (u *UnitController) Activate(e echo.Context) error {

	var unit dto.Unit

	if e.Bind(&unit) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := u.UnitUsecase.Activate(unit)

	if !unit.IsActived {
		return CheckResponse(e, err, err_code, config.DeactivateUnitSuccess)
	}

	return CheckResponse(e, err, err_code, config.ActivateUnitSuccess)
}

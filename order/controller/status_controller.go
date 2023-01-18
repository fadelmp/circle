package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type StatusController struct {
	StatusUsecase usecase.StatusUsecase
}

func ProviderStatusController(s usecase.StatusUsecase) StatusController {
	return StatusController{
		StatusUsecase: s,
	}
}

func (s *StatusController) GetAll(e echo.Context) error {

	statuses := s.StatusUsecase.GetAll()

	if len(statuses) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, statuses, config.GetStatusSuccess)
}

func (s *StatusController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	status := s.StatusUsecase.GetByID(uint(id))

	if status.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, status, config.GetStatusSuccess)
}

func (s *StatusController) Create(e echo.Context) error {

	var status dto.Status

	if e.Bind(&status) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := s.StatusUsecase.Create(status)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.CreateStatusSuccess)
}

func (s *StatusController) Update(e echo.Context) error {

	var status dto.Status

	if e.Bind(&status) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := s.StatusUsecase.Update(status)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.UpdateStatusSuccess)
}

func (s *StatusController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = s.StatusUsecase.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.DeleteStatusSuccess)
}

func (s *StatusController) Activate(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = s.StatusUsecase.ActiveStatus(uint(id), true)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.ActivateStatusSuccess)
}

func (s *StatusController) Deactivate(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = s.StatusUsecase.ActiveStatus(uint(id), false)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.DeactivateStatusSuccess)
}

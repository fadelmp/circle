package controller

import (
	"net/http"
	"order/config"
	"order/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
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
		return SuccessResponse(e, nil, config.StatusNotFound)
	}

	return SuccessResponse(e, statuses, config.GetStatusSuccess)
}

func (s *StatusController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	status := s.StatusUsecase.GetByID(uint(id))

	if status.ID == 0 {
		return SuccessResponse(e, nil, config.StatusNotFound)
	}

	return SuccessResponse(e, status, config.GetStatusSuccess)
}

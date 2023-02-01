package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type OrderUnitController struct {
	OrderUnitUsecase usecase.OrderUnitUsecase
}

func ProviderOrderUnitController(s usecase.OrderUnitUsecase) OrderUnitController {
	return OrderUnitController{
		OrderUnitUsecase: s,
	}
}

func (s *OrderUnitController) GetOrderUnits(e echo.Context) error {

	res := s.OrderUnitUsecase.GetOrderUnits()

	return CheckResponse(e, res)
}

func (s *OrderUnitController) GetOrderUnitById(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderUnitUsecase.GetOrderUnitById(uint(id))

	return CheckResponse(e, res)
}

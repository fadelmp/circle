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

	filter := e.QueryParam("filter")
	status := e.QueryParam("status")

	res := s.OrderUnitUsecase.GetOrderUnits(filter, status)

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

func (s *OrderUnitController) CreateOrderUnit(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := s.OrderUnitUsecase.CreateOrderUnit(request)

	return CheckResponse(e, res)
}

func (s *OrderUnitController) UpdateOrderUnit(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := s.OrderUnitUsecase.UpdateOrderUnit(request)

	return CheckResponse(e, res)

}

func (s *OrderUnitController) DeleteOrderUnit(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderUnitUsecase.DeleteOrderUnit(uint(id))

	return CheckResponse(e, res)
}

func (s *OrderUnitController) ActivateOrderUnit(e echo.Context) error {

	status := e.Param("status")
	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderUnitUsecase.ActivateOrderUnit(uint(id), status)

	return CheckResponse(e, res)
}

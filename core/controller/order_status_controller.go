package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type OrderStatusController struct {
	OrderStatusUsecase usecase.OrderStatusUsecase
}

func ProviderOrderStatusController(s usecase.OrderStatusUsecase) OrderStatusController {
	return OrderStatusController{
		OrderStatusUsecase: s,
	}
}

func (s *OrderStatusController) GetOrderStatuses(e echo.Context) error {

	res := s.OrderStatusUsecase.GetOrderStatuses()

	return CheckResponse(e, res)
}

func (s *OrderStatusController) GetOrderStatusById(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.GetOrderStatusById(uint(id))

	return CheckResponse(e, res)
}

func (s *OrderStatusController) CreateOrderStatus(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.CreateOrderStatus(request)

	return CheckResponse(e, res)
}

func (s *OrderStatusController) UpdateOrderStatus(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.UpdateOrderStatus(request)

	return CheckResponse(e, res)

}

func (s *OrderStatusController) DeleteOrderStatus(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.DeleteOrderStatus(uint(id))

	return CheckResponse(e, res)
}

func (s *OrderStatusController) ActivateOrderStatus(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.ActivateOrderStatus(uint(id))

	return CheckResponse(e, res)
}

func (s *OrderStatusController) DeactivateOrderStatus(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := s.OrderStatusUsecase.DeactivateOrderStatus(uint(id))

	return CheckResponse(e, res)
}

package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/usecase"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	OrderUsecase usecase.OrderUsecase
}

func ProviderOrderController(o usecase.OrderUsecase) OrderController {
	return OrderController{
		OrderUsecase: o,
	}
}

func (o *OrderController) GetAll(e echo.Context) error {

	var query dto.QueryParam

	if e.Bind(&query) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	orders := o.OrderUsecase.GetAll(query)

	if len(orders) == 0 {
		return SuccessResponse(e, nil, config.OrderNotFound)
	}

	return SuccessResponse(e, orders, config.GetOrderSuccess)
}

func (o *OrderController) GetByNumber(e echo.Context) error {

	order_number := e.Param("order_number")

	order := o.OrderUsecase.GetByNumber(order_number)

	if order.ID == 0 {
		return SuccessResponse(e, nil, config.GetOrderSuccess)
	}

	return SuccessResponse(e, order, config.GetOrderSuccess)
}

func (o *OrderController) Create(e echo.Context) error {

	var order dto.Order

	if e.Bind(&order) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err := o.OrderUsecase.Create(order)

	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, err.Error())
	}

	return SuccessResponse(e, nil, config.CreateOrderSuccess)
}

func (o *OrderController) Update(e echo.Context) error {

	var order dto.Order

	if e.Bind(&order) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err := o.OrderUsecase.Update(order)

	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, err.Error())
	}

	return SuccessResponse(e, nil, config.UpdateOrderSuccess)
}

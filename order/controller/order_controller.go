package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/usecase"
	"strconv"

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

	orders := o.OrderUsecase.GetAll()

	if len(orders) == 0 {
		return SuccessResponse(e, nil, config.OrderNotFound)
	}

	return SuccessResponse(e, orders, config.GetOrderSuccess)
}

func (o *OrderController) GetByCustomerID(e echo.Context) error {

	customer_id, _ := strconv.ParseUint(e.Param("customer_id"), 10, 32)

	orders := o.OrderUsecase.GetByCustomerID(uint(customer_id))

	if len(orders) == 0 {
		return SuccessResponse(e, nil, config.OrderNotFound)
	}

	return SuccessResponse(e, orders, config.GetOrderSuccess)

}

func (o *OrderController) GetByOrderNumber(e echo.Context) error {

	order_number := e.Param("order_number")

	order := o.OrderUsecase.GetByOrderNumber(order_number)

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

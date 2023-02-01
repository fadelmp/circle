package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type OrderController struct {
	OrderUsecase usecase.OrderUsecase
}

func ProviderOrderController(o usecase.OrderUsecase) OrderController {
	return OrderController{
		OrderUsecase: o,
	}
}

func (o *OrderController) GetOrders(e echo.Context) error {

	res := o.OrderUsecase.GetOrders()

	return CheckResponse(e, res)
}

func (o *OrderController) GetOrderByNumber(e echo.Context) error {

	number := e.Param("number")

	res := o.OrderUsecase.GetOrderByOrderNumber(number)

	return CheckResponse(e, res)
}

func (o *OrderController) GetOrderByCustomer(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("customer_id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := o.OrderUsecase.GetOrderByCustomer(uint(id))

	return CheckResponse(e, res)
}

func (o *OrderController) CreateOrder(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := o.OrderUsecase.CreateOrder(request)

	return CheckResponse(e, res)
}

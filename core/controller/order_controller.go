package controller

import (
	"core/config"
	"core/usecase"
	"net/http"

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

func (o *OrderController) CreateOrder(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := o.OrderUsecase.CreateOrder(request)

	return CheckResponse(e, res)
}

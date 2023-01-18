package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/usecase"

	"github.com/labstack/echo"
)

type OrderController struct {
	OrderUsecase usecase.OrderUsecase
}

func ProviderOrderController(o usecase.OrderUsecase) OrderController {
	return OrderController{
		OrderUsecase: o,
	}
}

func (o *OrderController) Create(e echo.Context) error {

	var order dto.Order

	if e.Bind(&order) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := o.OrderUsecase.Create(order)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.CreateOrderSuccess)
}

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

func ProviderOrderController(s usecase.OrderUsecase) OrderController {
	return OrderController{
		OrderUsecase: s,
	}
}

func (s *OrderController) CreateOrder(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res := s.OrderUsecase.CreateOrder(request)

	return CheckResponse(e, res)
}

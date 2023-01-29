package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type OrderArticleController struct {
	OrderArticleUsecase usecase.OrderArticleUsecase
}

func ProviderOrderArticleController(s usecase.OrderArticleUsecase) OrderArticleController {
	return OrderArticleController{
		OrderArticleUsecase: s,
	}
}

func (a *OrderArticleController) GetOrderArticles(e echo.Context) error {

	res := a.OrderArticleUsecase.GetOrderArticles()

	return CheckResponse(e, res)
}

func (a *OrderArticleController) GetOrderArticleById(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := a.OrderArticleUsecase.GetOrderArticleByID(uint(id))

	return CheckResponse(e, res)
}

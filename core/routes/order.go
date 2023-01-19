package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.POST("", api.CreateOrder)
	}

}

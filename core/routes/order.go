package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.GET("", api.GetOrders)
		order.GET("/number/:number", api.GetOrderByNumber)

		order.POST("", api.CreateOrder)
		order.PUT("", api.UpdateOrder)
	}

}

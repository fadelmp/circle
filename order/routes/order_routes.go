package routes

import (
	"order/controller"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.GET("", api.GetAll)
		order.GET("/number/:order_number", api.GetByNumber)

		order.POST("", api.Create)
		order.PUT("", api.Update)
	}
}

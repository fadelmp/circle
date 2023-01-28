package routes

import (
	"order/controller"

	"github.com/labstack/echo"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.GET("", api.GetAll)
		order.GET("/customer/:customer_id", api.GetByCustomerID)

		order.POST("", api.Create)
	}
}

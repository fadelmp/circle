package routes

import (
	"order/controller"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.GET("", api.GetAll)
		order.GET("/number/:order_number", api.GetByOrderNumber)
		order.GET("/customer/:customer_id", api.GetByCustomerID)
		order.GET("/status/:status_id", api.GetByStatusID)

		order.POST("", api.Create)
	}
}

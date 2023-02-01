package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func OrderStatusRoutes(routes *echo.Echo, api controller.OrderStatusController) {

	order_status := routes.Group("/order_status")
	{
		order_status.GET("", api.GetOrderStatuses)
		order_status.GET("/:id", api.GetOrderStatusById)
	}

}

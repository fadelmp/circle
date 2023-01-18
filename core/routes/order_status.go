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

		order_status.POST("", api.CreateOrderStatus)
		order_status.PUT("", api.UpdateOrderStatus)
		order_status.DELETE("/:id", api.DeleteOrderStatus)

		order_status.PATCH("/activate/:id", api.ActivateOrderStatus)
		order_status.PATCH("/deactivate/:id", api.DeactivateOrderStatus)
	}

}

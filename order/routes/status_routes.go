package routes

import (
	"order/controller"

	"github.com/labstack/echo/v4"
)

func StatusRoutes(routes *echo.Echo, api controller.StatusController) {

	status := routes.Group("/order_status")
	{
		status.GET("", api.GetAll)
		status.GET("/:id", api.GetByID)
	}
}

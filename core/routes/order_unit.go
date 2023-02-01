package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func OrderUnitRoutes(routes *echo.Echo, api controller.OrderUnitController) {

	order_unit := routes.Group("/order_unit")
	{
		order_unit.GET("", api.GetOrderUnits)
		order_unit.GET("/:id", api.GetOrderUnitById)
	}

}

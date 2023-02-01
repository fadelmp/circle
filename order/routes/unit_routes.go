package routes

import (
	"order/controller"

	"github.com/labstack/echo/v4"
)

func UnitRoutes(routes *echo.Echo, api controller.UnitController) {

	Unit := routes.Group("/order_unit")
	{
		Unit.GET("", api.GetAll)
		Unit.GET("/:id", api.GetByID)
	}
}

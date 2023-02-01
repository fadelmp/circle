package routes

import (
	"order/controller"

	"github.com/labstack/echo/v4"
)

func UnitRoutes(routes *echo.Echo, api controller.UnitController) {

	Unit := routes.Group("/Unit")
	{
		Unit.GET("", api.GetAll)
		Unit.GET("/:id", api.GetByID)

		Unit.POST("", api.Create)
		Unit.PUT("", api.Update)
		Unit.DELETE("/:id", api.Delete)

		Unit.PATCH("/:id/:is_actived", api.Activate)
	}
}

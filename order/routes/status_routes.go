package routes

import (
	"order/controller"

	"github.com/labstack/echo"
)

func StatusRoutes(routes *echo.Echo, api controller.StatusController) {

	status := routes.Group("/status")
	{
		status.GET("", api.GetAll)
		status.GET("/:id", api.GetByID)

		status.POST("", api.Create)
		status.PUT("", api.Update)
		status.DELETE("/:id", api.Delete)

		status.PATCH("/activate/:id", api.Activate)
		status.PATCH("/deactivate/:id", api.Deactivate)
	}
}

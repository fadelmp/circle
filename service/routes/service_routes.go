package routes

import (
	"service/controller"

	"github.com/labstack/echo/v4"
)

func ServiceRoutes(routes *echo.Echo, api controller.ServiceController) {

	service := routes.Group("/service")
	{
		service.GET("", api.GetAll)
		service.GET("/:id", api.GetByID)

		service.POST("", api.Create)
		service.PUT("", api.Update)
		service.DELETE("/:id", api.Delete)

		service.PATCH("/:id/:is_actived", api.Activate)
	}
}

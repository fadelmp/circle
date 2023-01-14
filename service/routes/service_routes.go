package routes

import (
	"service/controller"

	"github.com/labstack/echo"
)

func ServiceRoutes(routes *echo.Echo, api controller.ServiceController) {

	service := routes.Group("/service")
	{
		service.GET("/:id", api.GetByID)
		service.GET("", api.GetAll)

		service.POST("", api.Create)
		service.PUT("", api.Update)
		service.DELETE("/:ID", api.Delete)
	}
}

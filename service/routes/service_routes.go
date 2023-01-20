package routes

import (
	"service/controller"

	"github.com/labstack/echo"
)

func ServiceRoutes(routes *echo.Echo, api controller.ServiceController) {

	service := routes.Group("/service")
	{
		service.GET("", api.GetAll)
		service.GET("/active", api.GetActive)
		service.GET("/active/:ID", api.GetByID)

		service.POST("", api.Create)
		service.PUT("", api.Update)
		service.DELETE("/:ID", api.Delete)

		service.PATCH("/deactivate/:ID", api.Deactivate)
		service.PATCH("/activate/:ID", api.Activate)
	}
}

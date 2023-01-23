package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func ServiceRoutes(routes *echo.Echo, api controller.ServiceController) {

	service := routes.Group("/service")
	{
		service.GET("", api.GetServices)
		service.GET("/:id", api.GetServiceById)

		service.POST("", api.CreateService)
		service.PUT("", api.UpdateService)
		service.DELETE("/:id", api.DeleteService)

		service.PATCH("/:status/:id", api.ActivateService)
	}

}

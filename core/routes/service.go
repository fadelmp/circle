package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func ServiceRoutes(routes *echo.Echo, api controller.ServiceController) {

	service := routes.Group("/service")
	{
		service.GET("", api.GetServices)
		service.GET("/active", api.GetActiveServices)
		service.GET("/active/:id", api.GetServiceById)
		service.GET("/available", api.GetAvailableServices)

		service.POST("", api.CreateService)
		service.PUT("", api.UpdateService)
		service.DELETE("/:id", api.DeleteService)

		service.PATCH("/activate/:id", api.ActivateService)
		service.PATCH("/deactivate/:id", api.DeactivateService)
	}

}

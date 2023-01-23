package routes

import (
	"customer/controller"

	"github.com/labstack/echo"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.GET("", api.GetAll)
		customer.GET("/active", api.GetActive)
		customer.GET("/active/:ID", api.GetByID)
		customer.GET("/available", api.GetAvailable)

		customer.POST("", api.Create)
		customer.PUT("", api.Update)
		customer.DELETE("/:ID", api.Delete)

		customer.PATCH("/activate/:ID", api.Activate)
		customer.PATCH("/deactivate/:ID", api.Deactivate)
	}
}

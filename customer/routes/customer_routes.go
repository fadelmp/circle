package routes

import (
	"customer/controller"

	"github.com/labstack/echo"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.GET("", api.GetAll)
		customer.GET("/:ID", api.GetByID)

		customer.POST("", api.Create)
		customer.PUT("", api.Update)
		customer.DELETE("/:ID", api.Delete)

		customer.PATCH("/:Status/:ID", api.Activate)
	}
}

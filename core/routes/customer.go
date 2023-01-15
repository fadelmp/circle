package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.POST("", api.Create)
	}

}

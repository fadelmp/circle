package routes

import (
	"customer/controller"

	"github.com/labstack/echo/v4"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.GET("", api.GetAll)
		customer.GET("/:id", api.GetByID)

		customer.POST("", api.Create)
		customer.PUT("", api.Update)
		customer.DELETE("/:id", api.Delete)

		customer.PATCH("/:id/:is_actived", api.Activate)
	}
}

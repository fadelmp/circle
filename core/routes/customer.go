package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.GET("", api.GetCustomers)
		customer.GET("/:id", api.GetCustomerById)

		customer.POST("", api.CreateCustomer)
		customer.PUT("", api.UpdateCustomer)
		customer.DELETE("/:id", api.DeleteCustomer)

		customer.PATCH("/:status/:id", api.ActivateCustomer)
	}

}

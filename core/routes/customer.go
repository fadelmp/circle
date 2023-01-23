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

		customer.GET("", api.GetCustomers)
		customer.GET("/active", api.GetActiveCustomers)
		customer.GET("/active/:id", api.GetCustomerById)
		customer.GET("/available", api.GetAvailableCustomers)

		customer.POST("", api.CreateCustomer)
		customer.PUT("", api.UpdateCustomer)
		customer.DELETE("/:id", api.DeleteCustomer)

		customer.PATCH("/activate/:id", api.ActivateCustomer)
		customer.PATCH("/deactivate/:id", api.DeactivateCustomer)
	}

}

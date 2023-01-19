package routes

import (
	"core/injection"

	"github.com/labstack/echo"
)

func Init(routes *echo.Echo) *echo.Echo {

	// Location Route & Injection
	location := injection.LocationInjection()
	LocationRoutes(routes, location)

	// Customer Route & Injection
	customer := injection.CustomerInjection()
	CustomerRoutes(routes, customer)

	// Service Route & Injection
	service := injection.ServiceInjection()
	ServiceRoutes(routes, service)

	// Order Status Route & Injection
	order_status := injection.OrderStatusInjection()
	OrderStatusRoutes(routes, order_status)

	// Order Route & Injection
	order := injection.OrderInjection()
	OrderRoutes(routes, order)

	return routes
}

package routes

import (
	"core/injection"

	"github.com/labstack/echo"
)

func Init(routes *echo.Echo) *echo.Echo {

	// location Route & Injection
	location := injection.LocationInjection()
	LocationRoutes(routes, location)

	// customer Route & Injection
	customer := injection.CustomerInjection()
	CustomerRoutes(routes, customer)

	// service Route & Injection
	service := injection.ServiceInjection()
	ServiceRoutes(routes, service)

	return routes
}

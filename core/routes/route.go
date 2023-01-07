package routes

import (
	"core/injection"

	"github.com/labstack/echo"
)

func Init(routes *echo.Echo) *echo.Echo {

	// location Route & Injection
	location := injection.LocationInjection()
	LocationRoutes(routes, location)

	return routes
}

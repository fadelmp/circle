package routes

import (
	"customer/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// customer Route & Injection
	customer := injection.CustomerInjection(db)
	CustomerRoutes(routes, customer)

	return routes
}

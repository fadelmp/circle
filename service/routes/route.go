package routes

import (
	"service/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// Service Route & Injection
	service := injection.ServiceInjection(db)
	ServiceRoutes(routes, service)

	return routes
}

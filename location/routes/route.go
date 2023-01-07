package routes

import (
	"location/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// country Route & Injection
	country := injection.CountryInjection(db)
	CountryRoutes(routes, country)

	// province Route & Injection
	province := injection.ProvinceInjection(db)
	ProvinceRoutes(routes, province)

	// city Route & Injection
	city := injection.CityInjection(db)
	CityRoutes(routes, city)

	return routes
}

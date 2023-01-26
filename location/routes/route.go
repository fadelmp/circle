package routes

import (
	"location/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// country Route & Injection
	country := injection.CountryInjection(db, redis)
	CountryRoutes(routes, country)

	// province Route & Injection
	province := injection.ProvinceInjection(db, redis)
	ProvinceRoutes(routes, province)

	// city Route & Injection
	city := injection.CityInjection(db, redis)
	CityRoutes(routes, city)

	// district Route & Injection
	district := injection.DistrictInjection(db, redis)
	DistrictRoutes(routes, district)

	// sub district Route & Injection
	sub_district := injection.SubDistrictInjection(db, redis)
	SubDistrictRoutes(routes, sub_district)

	return routes
}

package routes

import (
	"customer/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// customer Route & Injection
	customer := injection.CustomerInjection(db, redis)
	CustomerRoutes(routes, customer)

	return routes
}

package routes

import (
	"service/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// Service Route & Injection
	service := injection.ServiceInjection(db, redis)
	ServiceRoutes(routes, service)

	return routes
}

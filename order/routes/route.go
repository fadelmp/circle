package routes

import (
	"order/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// Status Route & Injection
	status := injection.StatusInjection(db)
	StatusRoutes(routes, status)

	// Unit Route & Injection
	unit := injection.UnitInjection(db, redis)
	UnitRoutes(routes, unit)

	// Order Route & Injection
	order := injection.OrderInjection(db)
	OrderRoutes(routes, order)

	// Article Route & Injection
	article := injection.ArticleInjection(db)
	ArticleRoutes(routes, article)

	return routes
}

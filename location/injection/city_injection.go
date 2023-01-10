package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func CityInjection(db *gorm.DB, redis *redis.Client) controller.CityController {

	CityRepository := repository.ProviderCityRepository(db, redis)
	CityService := service.ProviderCityService(CityRepository)
	CityController := controller.ProviderCityController(CityService)

	return CityController
}

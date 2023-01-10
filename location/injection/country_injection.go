package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func CountryInjection(db *gorm.DB, redis *redis.Client) controller.CountryController {

	CountryRepository := repository.ProviderCountryRepository(db, redis)
	CountryService := service.ProviderCountryService(CountryRepository)
	CountryController := controller.ProviderCountryController(CountryService)

	return CountryController
}

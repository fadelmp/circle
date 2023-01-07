package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/jinzhu/gorm"
)

func CityInjection(db *gorm.DB) controller.CityController {

	CityRepository := repository.ProviderCityRepository(db)
	CityService := service.ProviderCityService(CityRepository)
	CityController := controller.ProviderCityController(CityService)

	return CityController
}

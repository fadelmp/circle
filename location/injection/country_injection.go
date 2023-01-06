package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/jinzhu/gorm"
)

func CountryInjection(db *gorm.DB) controller.CountryController {

	CountryRepository := repository.ProviderCountryRepository(db)
	CountryService := service.ProviderCountryService(CountryRepository)
	CountryController := controller.ProviderCountryController(CountryService)

	return CountryController
}

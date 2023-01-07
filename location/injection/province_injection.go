package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/jinzhu/gorm"
)

func ProvinceInjection(db *gorm.DB) controller.ProvinceController {

	ProvinceRepository := repository.ProviderProvinceRepository(db)
	ProvinceService := service.ProviderProvinceService(ProvinceRepository)
	ProvinceController := controller.ProviderProvinceController(ProvinceService)

	return ProvinceController
}

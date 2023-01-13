package injection

import (
	"service/controller"
	repository "service/repository"
	"service/usecase"

	"github.com/jinzhu/gorm"
)

func ServiceInjection(db *gorm.DB) controller.ServiceController {

	ServiceRepository := repository.ProviderServiceRepository(db)
	ServiceUsecase := usecase.ProviderServiceUsecase(ServiceRepository)
	ServiceController := controller.ProviderServiceController(ServiceUsecase)

	return ServiceController
}

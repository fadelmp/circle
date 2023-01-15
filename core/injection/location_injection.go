package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func LocationInjection() controller.LocationController {

	GetRequest := request.ProviderGetRequest()
	LocationUsecase := usecase.ProviderLocationUsecase(GetRequest)
	LocationController := controller.ProviderLocationController(LocationUsecase)

	return LocationController
}

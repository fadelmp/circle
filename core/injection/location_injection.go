package injection

import (
	"core/controller"
	"core/request"
	"core/service"
)

func LocationInjection() controller.LocationController {

	GetRequest := request.ProviderGetRequest()
	LocationService := service.ProviderLocationService(GetRequest)
	LocationController := controller.ProviderLocationController(LocationService)

	return LocationController
}

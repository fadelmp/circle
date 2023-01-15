package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func CustomerInjection() controller.CustomerController {

	GetRequest := request.ProviderGetRequest()
	PostRequest := request.ProviderPostRequest()

	CustomerUsecase := usecase.ProviderCustomerUsecase(GetRequest, PostRequest)
	CustomerController := controller.ProviderCustomerController(CustomerUsecase)

	return CustomerController
}

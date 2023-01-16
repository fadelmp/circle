package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func CustomerInjection() controller.CustomerController {

	GetRequest := request.ProviderGetRequest()
	PutRequest := request.ProviderPutRequest()
	PostRequest := request.ProviderPostRequest()
	PatchRequest := request.ProviderPatchRequest()
	DeleteRequest := request.ProviderDeleteRequest()

	CustomerUsecase := usecase.ProviderCustomerUsecase(
		GetRequest,
		PutRequest,
		PostRequest,
		PatchRequest,
		DeleteRequest,
	)

	CustomerController := controller.ProviderCustomerController(CustomerUsecase)

	return CustomerController
}

package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func ServiceInjection() controller.ServiceController {

	GetRequest := request.ProviderGetRequest()
	PutRequest := request.ProviderPutRequest()
	PostRequest := request.ProviderPostRequest()
	PatchRequest := request.ProviderPatchRequest()
	DeleteRequest := request.ProviderDeleteRequest()

	ServiceUsecase := usecase.ProviderServiceUsecase(
		GetRequest,
		PutRequest,
		PostRequest,
		PatchRequest,
		DeleteRequest,
	)

	ServiceController := controller.ProviderServiceController(ServiceUsecase)

	return ServiceController
}

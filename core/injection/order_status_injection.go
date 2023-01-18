package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderStatusInjection() controller.OrderStatusController {

	GetRequest := request.ProviderGetRequest()
	PutRequest := request.ProviderPutRequest()
	PostRequest := request.ProviderPostRequest()
	PatchRequest := request.ProviderPatchRequest()
	DeleteRequest := request.ProviderDeleteRequest()

	OrderStatusUsecase := usecase.ProviderOrderStatusUsecase(
		GetRequest,
		PutRequest,
		PostRequest,
		PatchRequest,
		DeleteRequest,
	)

	OrderStatusController := controller.ProviderOrderStatusController(OrderStatusUsecase)

	return OrderStatusController
}

package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderInjection() controller.OrderController {

	GetRequest := request.ProviderGetRequest()
	PutRequest := request.ProviderPutRequest()
	PostRequest := request.ProviderPostRequest()
	PatchRequest := request.ProviderPatchRequest()

	OrderUsecase := usecase.ProviderOrderUsecase(
		GetRequest,
		PutRequest,
		PostRequest,
		PatchRequest,
	)

	OrderController := controller.ProviderOrderController(OrderUsecase)

	return OrderController
}

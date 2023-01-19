package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderInjection() controller.OrderController {

	GetRequest := request.ProviderGetRequest()
	PostRequest := request.ProviderPostRequest()
	PatchRequest := request.ProviderPatchRequest()

	OrderUsecase := usecase.ProviderOrderUsecase(
		GetRequest,
		PostRequest,
		PatchRequest,
	)

	OrderController := controller.ProviderOrderController(OrderUsecase)

	return OrderController
}

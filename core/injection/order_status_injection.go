package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderStatusInjection() controller.OrderStatusController {

	GetRequest := request.ProviderGetRequest()
	OrderStatusUsecase := usecase.ProviderOrderStatusUsecase(GetRequest)
	OrderStatusController := controller.ProviderOrderStatusController(OrderStatusUsecase)

	return OrderStatusController
}

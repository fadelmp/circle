package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderUnitInjection() controller.OrderUnitController {

	GetRequest := request.ProviderGetRequest()
	OrderUnitUsecase := usecase.ProviderOrderUnitUsecase(GetRequest)
	OrderUnitController := controller.ProviderOrderUnitController(OrderUnitUsecase)

	return OrderUnitController
}

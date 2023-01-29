package injection

import (
	"core/controller"
	"core/request"
	"core/usecase"
)

func OrderArticleInjection() controller.OrderArticleController {

	GetRequest := request.ProviderGetRequest()
	OrderArticleUsecase := usecase.ProviderOrderArticleUsecase(GetRequest)
	OrderArticleController := controller.ProviderOrderArticleController(OrderArticleUsecase)

	return OrderArticleController
}

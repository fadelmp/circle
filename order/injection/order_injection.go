package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/jinzhu/gorm"
)

func OrderInjection(db *gorm.DB) controller.OrderController {

	OrderRepository := repository.ProviderOrderRepository(db)
	ServiceRepository := repository.ProviderServiceRepository(db)
	ArticleRepository := repository.ProviderArticleRepository(db)

	ServiceUsecase := usecase.ProviderServiceUsecase(ServiceRepository)
	ArticleUsecase := usecase.ProviderArticleUsecase(ArticleRepository, ServiceUsecase)
	OrderUsecase := usecase.ProviderOrderUsecase(OrderRepository, ArticleUsecase)

	OrderController := controller.ProviderOrderController(OrderUsecase)

	return OrderController
}

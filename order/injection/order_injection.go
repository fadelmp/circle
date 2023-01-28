package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/jinzhu/gorm"
)

func OrderInjection(db *gorm.DB) controller.OrderController {

	OrderRepository := repository.ProviderOrderRepository(db)
	OrderUsecase := usecase.ProviderOrderUsecase(OrderRepository)
	OrderController := controller.ProviderOrderController(OrderUsecase)

	return OrderController
}

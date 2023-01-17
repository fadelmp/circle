package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/jinzhu/gorm"
)

func StatusInjection(db *gorm.DB) controller.StatusController {

	StatusRepository := repository.ProviderStatusRepository(db)
	StatusUsecase := usecase.ProviderStatusUsecase(StatusRepository)
	StatusController := controller.ProviderStatusController(StatusUsecase)

	return StatusController
}

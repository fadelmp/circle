package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func StatusInjection(db *gorm.DB, redis *redis.Client) controller.StatusController {

	StatusRepository := repository.ProviderStatusRepository(db, redis)
	StatusUsecase := usecase.ProviderStatusUsecase(StatusRepository)
	StatusController := controller.ProviderStatusController(StatusUsecase)

	return StatusController
}

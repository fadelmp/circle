package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func UnitInjection(db *gorm.DB, redis *redis.Client) controller.UnitController {

	UnitRepository := repository.ProviderUnitRepository(db, redis)
	UnitUsecase := usecase.ProviderUnitUsecase(UnitRepository)
	UnitController := controller.ProviderUnitController(UnitUsecase)

	return UnitController
}

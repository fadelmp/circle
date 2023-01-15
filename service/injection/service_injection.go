package injection

import (
	"service/controller"
	repository "service/repository"
	"service/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func ServiceInjection(db *gorm.DB, redis *redis.Client) controller.ServiceController {

	ServiceRepository := repository.ProviderServiceRepository(db, redis)
	ServiceUsecase := usecase.ProviderServiceUsecase(ServiceRepository)
	ServiceController := controller.ProviderServiceController(ServiceUsecase)

	return ServiceController
}

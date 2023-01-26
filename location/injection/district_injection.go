package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func DistrictInjection(db *gorm.DB, redis *redis.Client) controller.DistrictController {

	DistrictRepository := repository.ProviderDistrictRepository(db, redis)
	DistrictService := service.ProviderDistrictService(DistrictRepository)
	DistrictController := controller.ProviderDistrictController(DistrictService)

	return DistrictController
}

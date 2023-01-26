package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func SubDistrictInjection(db *gorm.DB, redis *redis.Client) controller.SubDistrictController {

	SubDistrictRepository := repository.ProviderSubDistrictRepository(db, redis)
	SubDistrictService := service.ProviderSubDistrictService(SubDistrictRepository)
	SubDistrictController := controller.ProviderSubDistrictController(SubDistrictService)

	return SubDistrictController
}

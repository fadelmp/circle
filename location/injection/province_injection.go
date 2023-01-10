package injection

import (
	"location/controller"
	repository "location/repository"
	"location/service"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func ProvinceInjection(db *gorm.DB, redis *redis.Client) controller.ProvinceController {

	ProvinceRepository := repository.ProviderProvinceRepository(db, redis)
	ProvinceService := service.ProviderProvinceService(ProvinceRepository)
	ProvinceController := controller.ProviderProvinceController(ProvinceService)

	return ProvinceController
}

package injection

import (
	"customer/controller"
	repository "customer/repository"
	"customer/request"
	"customer/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func CustomerInjection(db *gorm.DB, redis *redis.Client) controller.CustomerController {

	LocationReq := request.ProviderLocationRequest()
	CustomerRepo := repository.ProviderCustomerRepository(db, redis)

	LocationUsecase := usecase.ProviderLocationUsecase(LocationReq)
	AddressUsecase := usecase.ProviderAddressUsecase(LocationUsecase)
	CustomerUsecase := usecase.ProviderCustomerUsecase(CustomerRepo, AddressUsecase)

	CustomerController := controller.ProviderCustomerController(CustomerUsecase)

	return CustomerController
}

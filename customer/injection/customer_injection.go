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

	AddressRepo := repository.ProviderAddressRepository(db)
	CustomerRepo := repository.ProviderCustomerRepository(db, redis)

	AddressUsecase := usecase.ProviderAddressUsecase(AddressRepo)
	LocationUsecase := usecase.ProviderLocationUsecase(LocationReq)

	CustomerUsecase := usecase.ProviderCustomerUsecase(
		CustomerRepo,
		LocationUsecase,
		AddressUsecase,
	)
	CustomerController := controller.ProviderCustomerController(CustomerUsecase)

	return CustomerController
}

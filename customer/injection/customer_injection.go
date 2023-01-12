package injection

import (
	"customer/controller"
	repository "customer/repository"
	"customer/service"

	"github.com/jinzhu/gorm"
)

func CustomerInjection(db *gorm.DB) controller.CustomerController {

	CustomerRepo := repository.ProviderCustomerRepository(db)
	AddressRepo := repository.ProviderAddressRepository(db)
	ContactPeopleRepo := repository.ProviderContactPeopleRepository(db)

	AddressService := service.ProviderAddressService(AddressRepo)
	ContactPeopleService := service.ProviderContactPeopleService(ContactPeopleRepo)
	CustomerService := service.ProviderCustomerService(CustomerRepo, AddressService, ContactPeopleService)

	CustomerController := controller.ProviderCustomerController(CustomerService)

	return CustomerController
}

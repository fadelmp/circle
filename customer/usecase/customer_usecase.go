package usecase

import (
	"customer/config"
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
	"errors"
)

type CustomerUsecaseContract interface {
	GetAll(string, string) []dto.ShowCustomer
	GetByID(uint) dto.Customer

	Create(entity.Customer) error
	Update(entity.Customer) error
	Delete(uint) error
	Activate(uint, string) error
}

type CustomerUsecase struct {
	CustomerRepository repository.CustomerRepository
	LocationUsecase    LocationUsecase
	AddressUsecase     AddressUsecase
}

func ProviderCustomerUsecase(
	c repository.CustomerRepository,
	l LocationUsecase,
	a AddressUsecase,
) CustomerUsecase {
	return CustomerUsecase{
		CustomerRepository: c,
		LocationUsecase:    l,
		AddressUsecase:     a,
	}
}

// Implementation

func (c *CustomerUsecase) GetAll(filter string, status string) []dto.ShowCustomer {

	var customers []entity.Customer

	if filter != "" {
		customers = c.CustomerRepository.GetByFilter(filter)
	} else if status == "available" {
		customers = c.CustomerRepository.GetAvailable()
	} else if status == "active" {
		customers = c.CustomerRepository.GetActive()
	} else {
		customers = c.CustomerRepository.GetAll()
	}

	locations := c.LocationUsecase.CheckLocation(customers)

	return mapper.ToShowCustomerDtoList(customers, locations)
}

func (c *CustomerUsecase) GetByID(id uint) dto.Customer {

	customer := c.CustomerRepository.GetByID(id)

	return mapper.ToCustomerDto(customer)
}

func (c *CustomerUsecase) Create(dto dto.Customer) error {

	// change customer dto to entity to put on database
	customer_entity := mapper.ToCustomerEntity(dto)
	customer_entity.Base = entity.BaseCreate()

	// create customer data
	return c.CustomerRepository.Create(customer_entity)
}

func (c *CustomerUsecase) Update(dto dto.Customer) error {

	if !c.CheckID(dto.ID) {
		return errors.New(config.CustomerNotFound)
	}

	customer_entity := mapper.ToCustomerEntity(dto)
	customer_entity.Base = entity.BaseUpdate()

	// Update Customer Data
	err := c.CustomerRepository.Update(customer_entity)

	return err
}

func (c *CustomerUsecase) Delete(id uint) error {

	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	var customer_entity entity.Customer

	customer_entity.ID = id
	customer_entity.Base = entity.BaseDelete()

	return c.CustomerRepository.ChangeStatus(customer_entity)
}

func (c *CustomerUsecase) Activate(id uint, status string) error {

	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	is_active := true
	if status == "deactivate" {
		is_active = false
	}

	var customer_entity entity.Customer

	customer_entity.ID = id
	customer_entity.Base = entity.BaseActivate(is_active)

	return c.CustomerRepository.ChangeStatus(customer_entity)
}

func (c *CustomerUsecase) CheckID(id uint) bool {

	customer_data := c.CustomerRepository.GetByID(id)

	if customer_data.ID == 0 || customer_data.Is_Deleted {
		return false
	}

	return true
}

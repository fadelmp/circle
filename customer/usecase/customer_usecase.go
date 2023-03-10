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
	GetAll(dto.QueryParam) []dto.Customer
	GetByID(uint) dto.Customer

	Create(entity.Customer) (dto.Customer, error, int)
	Update(entity.Customer) (error, int)
	Delete(entity.Customer) (error, int)
	Activate(entity.Customer) (error, int)
}

type CustomerUsecase struct {
	CustomerRepository repository.CustomerRepository
	AddressUsecase     AddressUsecase
}

func ProviderCustomerUsecase(
	c repository.CustomerRepository,
	a AddressUsecase,
) CustomerUsecase {
	return CustomerUsecase{
		CustomerRepository: c,
		AddressUsecase:     a,
	}
}

// Implementation

func (c *CustomerUsecase) GetAll(dto dto.QueryParam) []dto.Customer {

	var customers []entity.Customer

	if dto.Filter != "" {
		customers = c.CustomerRepository.GetByFilter(dto.Filter)
	} else if dto.Name != "" {
		customers = c.CustomerRepository.GetByName(dto.Name)
	} else if dto.Status == "available" {
		customers = c.CustomerRepository.GetAvailable()
	} else if dto.Status == "active" {
		customers = c.CustomerRepository.GetActive()
	} else {
		customers = c.CustomerRepository.GetAll()
	}

	locations := c.AddressUsecase.CheckList(customers)

	return mapper.ToCustomerDtoList(customers, locations)
}

func (c *CustomerUsecase) GetByID(id uint) dto.Customer {

	customer := c.CustomerRepository.GetByID(id)

	location := c.AddressUsecase.Check(customer)

	return mapper.ToCustomerDto(customer, location)
}

func (c *CustomerUsecase) Create(dto dto.Customer) (dto.Customer, error, int) {

	// check phone whether customer exists
	if !c.CheckPhone(dto) {
		return dto, errors.New(config.CustomerExists), 2
	}

	// change customer dto to entity to put on database
	customer_entity := mapper.ToCustomerEntity(dto, entity.BaseCreate())

	// create customer data
	customer, err := c.CustomerRepository.Create(customer_entity)

	dto_customer := mapper.ToCustomerDto(customer, "")

	return dto_customer, err, 0
}

func (c *CustomerUsecase) Update(dto dto.Customer) (error, int) {

	if !c.CheckID(dto.ID) {
		return errors.New(config.CustomerNotFound), 1
	}

	// check phone whether customer exists
	if !c.CheckPhone(dto) {
		return errors.New(config.CustomerExists), 2
	}

	customer_entity := mapper.ToCustomerEntity(dto, entity.BaseUpdate())

	// Update Customer Data
	return c.CustomerRepository.Update(customer_entity), 0
}

func (c *CustomerUsecase) Delete(dto dto.Customer) (error, int) {

	if !c.CheckID(dto.ID) {
		return errors.New(config.CustomerNotFound), 1
	}

	customer_entity := mapper.ToCustomerEntity(dto, entity.BaseDelete())

	return c.CustomerRepository.ChangeStatus(customer_entity), 2
}

func (c *CustomerUsecase) Activate(dto dto.Customer) (error, int) {

	if !c.CheckID(dto.ID) {
		return errors.New(config.CustomerNotFound), 1
	}

	customer_entity := mapper.ToCustomerEntity(dto, entity.BaseActivate(dto.IsActived))

	return c.CustomerRepository.ChangeStatus(customer_entity), 0
}

func (c *CustomerUsecase) CheckPhone(dto dto.Customer) bool {

	customer := c.CustomerRepository.GetByPhone(dto.Phone)

	if customer.ID != 0 && customer.ID != dto.ID &&
		!customer.Base.Is_Deleted {
		return false
	}

	return true
}

func (c *CustomerUsecase) CheckID(id uint) bool {

	customer_data := c.CustomerRepository.GetByID(id)

	if customer_data.ID == 0 || customer_data.Is_Deleted {
		return false
	}

	return true
}

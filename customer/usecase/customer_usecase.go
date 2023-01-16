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
	GetAll() []dto.Customer
	GetByID(uint) dto.Customer

	Create(entity.Customer) error
	Update(entity.Customer) error
	Delete(entity.Customer) error
	ActiveStatus(entity.Customer) error
}

type CustomerUsecase struct {
	CustomerRepository repository.CustomerRepository
	LocationUsecase    LocationUsecase
	AddressUsecase     AddressUsecase
	CompanyUsecase     CompanyUsecase
}

func ProviderCustomerUsecase(
	c repository.CustomerRepository,
	l LocationUsecase,
	a AddressUsecase,
	cp CompanyUsecase,
) CustomerUsecase {
	return CustomerUsecase{
		CustomerRepository: c,
		LocationUsecase:    l,
		AddressUsecase:     a,
		CompanyUsecase:     cp,
	}
}

// Implementation

func (c *CustomerUsecase) GetAll() []dto.Customer {

	customers := c.CustomerRepository.GetAll()

	return mapper.ToCustomerDtoList(customers)
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
	customer, err := c.CustomerRepository.Create(customer_entity)
	if err != nil {
		return err
	}

	// create address data
	if err := c.AddressUsecase.Create(dto.Address, customer.ID); err != nil {
		return err
	}

	// create contact person data
	if err := c.CompanyUsecase.Create(dto.Company, customer.ID); err != nil {
		return err
	}

	return nil
}

func (c *CustomerUsecase) Update(dto dto.Customer) error {

	if !c.CheckID(dto.ID) {
		return errors.New(config.CustomerNotFound)
	}

	customer_entity := mapper.ToCustomerEntity(dto)
	customer_entity.Base = entity.BaseUpdate()

	// create customer data
	err := c.CustomerRepository.Update(customer_entity)
	if err != nil {
		return err
	}

	// create address data
	if err := c.AddressUsecase.Update(dto.Address, dto.ID); err != nil {
		return err
	}

	// create contact person data
	if err := c.CompanyUsecase.Update(dto.Company, dto.ID); err != nil {
		return err
	}

	return nil
}

func (c *CustomerUsecase) Delete(id uint) error {

	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	var customer_entity entity.Customer

	customer_entity.ID = id
	customer_entity.Base = entity.BaseDelete()

	return c.CustomerRepository.Delete(customer_entity)
}

func (c *CustomerUsecase) ActiveStatus(id uint, is_active bool) error {

	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	var customer_entity entity.Customer

	customer_entity.ID = id
	customer_entity.Base = entity.BaseActivate(is_active)

	return c.CustomerRepository.ActiveStatus(customer_entity)
}

func (c *CustomerUsecase) CheckID(id uint) bool {

	customer_data := c.CustomerRepository.GetByID(id)

	if customer_data.ID == 0 ||
		!customer_data.Is_Actived ||
		customer_data.Is_Deleted {
		return false
	}

	return true
}

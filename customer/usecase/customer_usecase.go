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

func (c *CustomerUsecase) Create(dto dto.Customer) error {

	// Check Phone number first before create data, return error if phone exists
	if !c.CheckName(dto) {
		return errors.New(config.CustomerExists)
	}

	// change customer dto to entity to put on database
	customer_entity := mapper.ToCustomerEntity(dto)

	// create customer data
	customer, err := c.CustomerRepository.Create(customer_entity)
	if err != nil {
		return err
	}

	// create address data
	if c.AddressUsecase.Create(dto.Address, customer.ID) != nil {
		return err
	}

	// create contact person data
	if c.CompanyUsecase.Create(dto.Company, customer.ID) != nil {
		return err
	}

	return nil
}

func (c *CustomerUsecase) CheckName(dto dto.Customer) bool {

	name := dto.Name

	// get customer data by phone number
	customer_data := c.CustomerRepository.GetByName(name)

	// return false if data id exists, and if is_active value is true
	if customer_data.ID != 0 &&
		customer_data.Is_Actived &&
		!customer_data.Is_Deleted {
		return false
	}

	return true
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

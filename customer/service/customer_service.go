package service

import (
	"customer/config"
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
	"errors"
)

type CustomerServiceContract interface {
	Create(entity.Customer) error
}

type CustomerService struct {
	CustomerRepository   repository.CustomerRepository
	AddressService       AddressService
	ContactPeopleService ContactPeopleService
}

func ProviderCustomerService(
	c repository.CustomerRepository,
	a AddressService,
	cp ContactPeopleService,
) CustomerService {
	return CustomerService{
		CustomerRepository:   c,
		AddressService:       a,
		ContactPeopleService: cp,
	}
}

// Implementation

func (c *CustomerService) Create(dto dto.Customer) error {

	// Check Phone number first before create data, return error if phone exists
	if !c.CheckName(dto) {
		return errors.New(config.CustomerExists)
	}

	// change customer dto to entity to put on database
	customer_entity := mapper.ToCustomerEntity(dto)

	// create customer data
	customer, err := c.CustomerRepository.Create(customer_entity)

	// return error if create customer data error
	if err != nil {
		return err
	}

	// create address data
	if c.AddAddress(customer.ID, dto.Addresses) != nil {
		return err
	}

	// create contact person data
	if c.AddContactPeople(customer.ID, dto.ContactPeoples) != nil {
		return err
	}

	return nil
}

func (c *CustomerService) CheckName(dto dto.Customer) bool {

	name := dto.Name

	// get customer data by phone number
	customer_data := c.CustomerRepository.GetByName(name)

	// return false if data id exists, and if is_active value is true
	if customer_data.ID != 0 && customer_data.IsActive {
		return false
	}

	return true
}

func (c *CustomerService) AddAddress(id uint, dto []dto.Address) error {

	for _, value := range dto {

		value.CustomerID = id

		// call create on address repository to create address value
		if err := c.AddressService.Create(value); err != nil {
			return err
		}
	}

	return nil
}

func (c *CustomerService) AddContactPeople(id uint, dto []dto.ContactPeople) error {

	for _, value := range dto {

		value.CustomerID = id

		if err := c.ContactPeopleService.Create(value); err != nil {
			return err
		}
	}

	return nil
}

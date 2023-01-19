package usecase

import (
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
)

type AddressUsecaseContract interface {
	Create(entity.Address, uint) error
	Update(entity.Address, uint) error
}

type AddressUsecase struct {
	AddressRepository repository.AddressRepository
}

func ProviderAddressUsecase(a repository.AddressRepository) AddressUsecase {
	return AddressUsecase{AddressRepository: a}
}

// Implementation

func (a *AddressUsecase) Create(dto dto.Address, customer_id uint) error {

	// Map dto to entity
	address_entity := mapper.ToAddressEntity(dto, customer_id)

	// Create address data
	err := a.AddressRepository.Create(address_entity)

	// Map entity to dto
	return err
}

func (a *AddressUsecase) Update(dto dto.Address, customer_id uint) error {

	// Map dto to entity
	address_entity := mapper.ToAddressEntity(dto, customer_id)

	// Update Address Data
	err := a.AddressRepository.Update(address_entity)

	return err
}

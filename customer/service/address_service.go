package service

import (
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
)

type AddressServiceContract interface {
	Create(entity.Address) error
}

type AddressService struct {
	AddressRepository repository.AddressRepository
}

func ProviderAddressService(a repository.AddressRepository) AddressService {
	return AddressService{AddressRepository: a}
}

// Implementation

func (a *AddressService) Create(dto dto.Address) error {

	// Map dto to entity
	address_entity := mapper.ToAddressEntity(dto)

	// Create address data
	_, err := a.AddressRepository.Create(address_entity)

	// Map entity to dto
	return err
}

package usecase

import (
	entity "customer/entity"
)

type AddressUsecaseContract interface {
	Create(entity.Address, uint) error
	Update(entity.Address, uint) error
}

type AddressUsecase struct{}

func ProviderAddressUsecase() AddressUsecase {
	return AddressUsecase{}
}

// Implementation

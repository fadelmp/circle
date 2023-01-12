package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type AddressRepositoryContract interface {
	Create(entity.Address) (entity.Address, error)
}

type AddressRepository struct {
	DB *gorm.DB
}

func ProviderAddressRepository(DB *gorm.DB) AddressRepository {
	return AddressRepository{DB: DB}
}

// Implementation

func (a *AddressRepository) Create(address entity.Address) (entity.Address, error) {

	// Create Address
	err := a.DB.Create(&address).Error

	return address, err
}

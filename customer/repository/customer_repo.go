package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type CustomerRepositoryContract interface {
	GetByName(string) entity.Customer
	Create(entity.Customer) (entity.Customer, error)
}

type CustomerRepository struct {
	DB *gorm.DB
}

func ProviderCustomerRepository(DB *gorm.DB) CustomerRepository {
	return CustomerRepository{DB: DB}
}

// Implementation

func (c *CustomerRepository) GetByName(name string) entity.Customer {

	var customer entity.Customer

	// Find All Province
	c.DB.Model(&customer).Where("name=?", name).Find(&customer)

	return customer
}

func (c *CustomerRepository) Create(customer entity.Customer) (entity.Customer, error) {

	// Create Customer Data
	err := c.DB.Create(&customer).Error

	return customer, err
}

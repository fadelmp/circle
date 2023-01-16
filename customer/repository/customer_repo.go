package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type CustomerRepositoryContract interface {
	GetAll() []entity.Customer
	GetByID(uint) entity.Customer
	GetByName(string) entity.Customer

	Create(entity.Customer) (entity.Customer, error)
	Update(entity.Customer) error
	Delete(entity.Customer) error
	ActiveStatus(entity.Customer) error
}

type CustomerRepository struct {
	DB *gorm.DB
}

func ProviderCustomerRepository(DB *gorm.DB) CustomerRepository {
	return CustomerRepository{DB: DB}
}

// Implementation

func (c *CustomerRepository) GetAll() []entity.Customer {

	var customers []entity.Customer
	var customer entity.Customer

	c.DB.Model(&customer).Preload("Address").Preload("Company").Find(&customers)

	return customers
}

func (c *CustomerRepository) GetByID(id uint) entity.Customer {

	var customer entity.Customer

	c.DB.Where("id=?", id).Preload("Address").Preload("Company").Find(&customer)

	return customer
}

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

func (c *CustomerRepository) Update(customer entity.Customer) error {

	// update Service by id
	err := c.DB.Model(&customer).Update(&customer).Error

	return err
}

func (c *CustomerRepository) Delete(customer entity.Customer) error {

	// delete Service by id, by change is active value to false
	err := c.DB.Model(&customer).Where("id=?", customer.ID).Updates(map[string]interface{}{
		"is_actived": customer.Base.Is_Actived,
		"is_deleted": customer.Base.Is_Deleted,
		"updated_at": customer.Base.Updated_At,
		"updated_by": customer.Base.Updated_By,
	}).Error

	return err
}

func (c *CustomerRepository) ActiveStatus(customer entity.Customer) error {

	// delete Service by id, by change is active value to false
	err := c.DB.Model(&customer).Where("id=?", customer.ID).Updates(map[string]interface{}{
		"is_actived": customer.Base.Is_Actived,
		"updated_at": customer.Base.Updated_At,
		"updated_by": customer.Base.Updated_By,
	}).Error

	return err
}

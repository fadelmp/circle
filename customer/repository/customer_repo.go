package repository

import (
	"customer/config"
	entity "customer/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type CustomerRepositoryContract interface {
	GetAll() []entity.Customer
	GetActive() []entity.Customer
	GetAvailable() []entity.Customer

	GetByID(uint) entity.Customer
	GetByName(string) entity.Customer
	GetByFilter(string) []entity.Customer

	Create(entity.Customer) (entity.Customer, error)
	Update(entity.Customer) error
	ChangeStatus(entity.Customer) error
}

type CustomerRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderCustomerRepository(DB *gorm.DB, Redis *redis.Client) CustomerRepository {
	return CustomerRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (c *CustomerRepository) GetAll() []entity.Customer {

	var customers []entity.Customer

	query := c.DB.Model(&entity.Customer{}).Preload("Address").Order("id asc").Find(&customers)
	keys := "customers"

	// Get Service All
	config.CheckRedisQuery(c.Redis, query, keys)

	return customers
}

func (c *CustomerRepository) GetActive() []entity.Customer {

	var customers []entity.Customer

	query := c.DB.Model(&entity.Customer{}).
		Where("customers.is_actived=?", true).
		Preload("Address").
		Order("id asc").Find(&customers)
	keys := "customers_active"

	// Get Service All
	config.CheckRedisQuery(c.Redis, query, keys)

	return customers
}

func (c *CustomerRepository) GetAvailable() []entity.Customer {

	var customers []entity.Customer

	query := c.DB.Model(&entity.Customer{}).
		Where("is_deleted=?", false).
		Preload("Address").Order("id asc").Find(&customers)
	keys := "customers_available"

	config.CheckRedisQuery(c.Redis, query, keys)

	return customers
}

func (c *CustomerRepository) GetByID(id uint) entity.Customer {

	var customer entity.Customer

	query := c.DB.Where("id=?", id).Preload("Address").Find(&customer)
	keys := "customer_id_" + strconv.FormatUint(uint64(id), 10)

	// Get Service By Id
	config.CheckRedisQuery(c.Redis, query, keys)

	return customer
}

func (c *CustomerRepository) GetByName(name string) entity.Customer {

	var customer entity.Customer

	query := c.DB.Where("name=?", name).Find(&customer)
	keys := "customer_name_" + name

	// Get Service by Name
	config.CheckRedisQuery(c.Redis, query, keys)

	return customer
}

func (c *CustomerRepository) GetByFilter(filter string) []entity.Customer {

	var customers []entity.Customer

	query := c.DB.Order("id asc").
		Where("is_deleted=?", false).
		Where("name LIKE ?", "%"+filter+"%").
		Preload("Address").Find(&customers)
	keys := "service_filter_" + filter

	// Get Service by Name
	config.CheckRedisQuery(c.Redis, query, keys)

	return customers
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

func (c *CustomerRepository) ChangeStatus(customer entity.Customer) error {

	// delete Service by id, by change is active value to false
	err := c.DB.Model(&customer).Where("id=?", customer.ID).Updates(map[string]interface{}{
		"is_actived": customer.Base.Is_Actived,
		"is_deleted": customer.Base.Is_Deleted,
		"updated_at": customer.Base.Updated_At,
		"updated_by": customer.Base.Updated_By,
	}).Error

	return err
}

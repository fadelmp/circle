package repository

import (
	entity "order/entity"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderRepositoryContract interface {
	GetAll() []entity.Order

	GetByFilter(string, uint, time.Time, time.Time)
	GetByOrderNumber(string) entity.Order
	GetByCustomerID(uint) []entity.Order
	GetByStatusID(uint) []entity.Order
	GetByDate(time.Time, time.Time) []entity.Order

	Create(entity.Order) error
	Update(entity.Order) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func ProviderOrderRepository(DB *gorm.DB) OrderRepository {
	return OrderRepository{DB: DB}
}

func (o *OrderRepository) GetAll() []entity.Order {

	var orders []entity.Order

	o.DB.Model(&entity.Order{}).Order("id asc").Preload("Status").
		Preload("Articles").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByFilter(
	search string,
	status_id uint,
	from time.Time,
	to time.Time,
) []entity.Order {

	var orders []entity.Order

	data := o.DB.Model(&entity.Order{})

	if search != "" {
		data = data.Where("number=?", search).Or("customer_name=?", search)
	}

	if status_id != 0 {
		data = data.Where("status_id=?", status_id)
	}

	if from.IsZero() && to.IsZero() {
		data = data.Where("created_at BETWEEN ? AND ?", from, to)
	}

	data.Order("id ASC").Preload("Status").Preload("Articles").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByOrderNumber(order_number string) entity.Order {

	var order entity.Order

	o.DB.Where("number=?", order_number).Preload("Status").
		Preload("Articles").Preload("Articles.Services").
		Preload("Articles.Unit").Preload("Articles.Status").
		Find(&order)

	return order
}

func (o *OrderRepository) GetByCustomerID(customer_id uint) []entity.Order {

	var orders []entity.Order

	o.DB.Where("customer_id=?", customer_id).Order("id asc").
		Preload("Articles").Preload("Status").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByStatusID(status_id uint) []entity.Order {

	var orders []entity.Order

	o.DB.Where("status_id=?", status_id).Order("id asc").
		Preload("Articles").Preload("Status").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByDate(first_date time.Time, last_date time.Time) []entity.Order {

	var orders []entity.Order

	o.DB.Where("created_at BETWEEN ? AND ?", first_date, last_date).Find(&orders)

	return orders
}

func (o *OrderRepository) Create(order entity.Order) error {

	return o.DB.Create(&order).Error
}

func (o *OrderRepository) Update(order entity.Order) error {

	// update Order by id
	return o.DB.Model(&order).Update(&order).Error
}
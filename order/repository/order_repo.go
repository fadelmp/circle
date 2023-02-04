package repository

import (
	entity "order/entity"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderRepositoryContract interface {
	GetAll() []entity.Order

	GetByOrderNumber(string) entity.Order
	GetByCustomerID(uint) []entity.Order
	GetByStatusID(uint) []entity.Order
	GetByDate(time.Time, time.Time) []entity.Order

	Create(entity.Order) (entity.Order, error)
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

func (o *OrderRepository) GetByOrderNumber(order_number string) entity.Order {

	var order entity.Order

	o.DB.Where("number=?", order_number).Preload("Status").
		Preload("Articles").Preload("Articles.Services").
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

func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {

	err := o.DB.Create(&order).Error

	return order, err
}

func (o *OrderRepository) Update(order entity.Order) error {

	// delete Service by id, by change is active value to false
	return o.DB.Model(&order).Where("id=?", order.ID).Updates(map[string]interface{}{
		"status_id":  order.StatusID,
		"updated_at": order.Base.Updated_At,
		"updated_by": order.Base.Updated_By,
	}).Error
}

package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type OrderRepositoryContract interface {
	GetAll() []entity.Order
	GetByID(uint) entity.Order
	GetByCustomerID(uint) []entity.Order

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

	o.DB.Model(&entity.Order{}).Order("id asc").Preload("Articles").Preload("Status").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByID(id uint) entity.Order {

	var order entity.Order

	o.DB.Where("id=?", id).Order("id asc").Preload("Articles").Preload("Status").Find(&order)

	return order
}

func (o *OrderRepository) GetByCustomerID(customer_id uint) []entity.Order {

	var orders []entity.Order

	o.DB.Where("customer_id=?", customer_id).Order("id asc").Preload("Articles").Preload("Status").Find(&orders)

	return orders
}

func (o *OrderRepository) Create(order entity.Order) error {

	return o.DB.Create(&order).Error
}

func (o *OrderRepository) Update(order entity.Order) error {

	// delete Service by id, by change is active value to false
	return o.DB.Model(&order).Where("id=?", order.ID).Updates(map[string]interface{}{
		"status_id":  order.StatusID,
		"updated_at": order.Base.Updated_At,
		"updated_by": order.Base.Updated_By,
	}).Error
}

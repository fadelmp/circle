package repository

import (
	entity "order/entity"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderRepositoryContract interface {
	GetByNumber(string) entity.Order
	Create(entity.Order) error
	Update(entity.Order) error
	ChangeStatus(entity.Order) error

	Filter() *gorm.DB
	FilterSearch(*gorm.DB, string) *gorm.DB
	FilterStatus(*gorm.DB, uint) *gorm.DB
	FilterCustomer(*gorm.DB, uint) *gorm.DB
	FilterTime(*gorm.DB, time.Time, time.Time) *gorm.DB
	Exec(*gorm.DB) []entity.Order
}

type OrderRepository struct {
	DB *gorm.DB
}

func ProviderOrderRepository(DB *gorm.DB) OrderRepository {
	return OrderRepository{DB: DB}
}

func (o *OrderRepository) GetAll() []entity.Order {

	var orders []entity.Order

	o.DB.Model(&entity.Order{}).Order("id asc").
		Preload("Status").Preload("Articles").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByNumber(order_number string) entity.Order {

	var order entity.Order

	o.DB.Where("number=?", order_number).Preload("Status").
		Preload("Articles").Preload("Articles.Services").
		Preload("Articles.Unit").Preload("Articles.Status").
		Find(&order)

	return order
}

func (o *OrderRepository) Create(order entity.Order) error {

	return o.DB.Create(&order).Error
}

func (o *OrderRepository) Update(order entity.Order) error {

	// update Order by id
	return o.DB.Model(&order).Update(&order).Error
}

func (o *OrderRepository) ChangeStatus(order entity.Order) error {

	// change is actived and is deleted value
	return o.DB.Model(&order).Where("id=?", order.ID).Updates(map[string]interface{}{
		"is_actived": order.Base.Is_Actived,
		"is_deleted": order.Base.Is_Deleted,
		"updated_at": order.Base.Updated_At,
		"updated_by": order.Base.Updated_By,
	}).Error
}

func (o *OrderRepository) Filter() *gorm.DB {

	return o.DB.Model(&entity.Order{})
}

func (o *OrderRepository) FilterSearch(data *gorm.DB, search string) *gorm.DB {

	return data.Where("number=?", search).
		Or("customer_name=?", search)
}

func (o *OrderRepository) FilterStatus(data *gorm.DB, status_id uint) *gorm.DB {

	return data.Where("status_id=?", status_id)
}

func (o *OrderRepository) FilterCustomer(data *gorm.DB, customer_id uint) *gorm.DB {

	return data.Where("customer_id=?", customer_id)
}

func (o *OrderRepository) FilterTime(data *gorm.DB, from time.Time, to time.Time) *gorm.DB {

	return data.Where("created_at BETWEEN ? AND ?", from, to)
}

func (o *OrderRepository) Exec(data *gorm.DB) []entity.Order {

	var orders []entity.Order

	data.Order("id asc").Preload("Status").Preload("Articles").Find(&orders)

	return orders
}

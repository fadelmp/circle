package usecase

import (
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
	"strconv"
	"time"
)

type OrderUsecaseContract interface {
	GetAll() []dto.ShowOrder
	GetByCustomer(uint) []dto.ShowOrder

	Create(entity.Order) error
	Update(entity.Order) error
}

type OrderUsecase struct {
	OrderRepository repository.OrderRepository
}

func ProviderOrderUsecase(o repository.OrderRepository) OrderUsecase {
	return OrderUsecase{
		OrderRepository: o,
	}
}

func (o *OrderUsecase) GetAll() []dto.ShowOrder {

	orders := o.OrderRepository.GetAll()

	return mapper.ToShowOrderDtoList(orders)
}

func (o *OrderUsecase) GetByCustomerID(customer_id uint) []dto.ShowOrder {

	orders := o.OrderRepository.GetByCustomerID(customer_id)

	return mapper.ToShowOrderDtoList(orders)
}

func (o *OrderUsecase) Create(dto dto.Order) error {

	dto.Number = o.GenerateOrderNumber()
	dto.StatusID = 1

	order_entity := mapper.ToOrderEntity(dto)
	order_entity.Base = entity.BaseCreate()

	return o.OrderRepository.Create(order_entity)
}

func (o *OrderUsecase) GenerateOrderNumber() string {

	t := time.Now()
	return strconv.FormatInt(t.UnixNano(), 10)

}

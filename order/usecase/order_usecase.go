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

	order_entity := mapper.ToOrderEntity(dto)
	order_entity.Base = entity.BaseCreate()

	return o.OrderRepository.Create(order_entity)
}

func (o *OrderUsecase) GenerateOrderNumber() string {

	date := GetDate()
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)

	orders := o.OrderRepository.GetByDate(Bod(today), Bod(tomorrow))
	count := len(orders) + 1
	number_str := strconv.FormatUint(uint64(count), 10)

	if len(number_str) == 1 {
		number_str = "000" + number_str
	} else if len(number_str) == 2 {
		number_str = "00" + number_str
	} else if (len(number_str)) == 3 {
		number_str = "0" + number_str
	} else {
		number_str = "" + number_str
	}

	return date + number_str
}

func GetDate() string {

	t := time.Now()

	day := t.Day()
	month := t.Month()
	year := t.Year() % 1e2

	day_str := strconv.FormatUint(uint64(day), 10)
	month_str := strconv.FormatUint(uint64(month), 10)
	year_str := strconv.FormatUint(uint64(year), 10)

	if month < 10 {
		month_str = "0" + strconv.FormatUint(uint64(month), 10)
	}

	return day_str + month_str + year_str
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

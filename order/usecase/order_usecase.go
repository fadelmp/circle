package usecase

import (
	"errors"
	"order/config"
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
	"strconv"
	"time"
)

type OrderUsecaseContract interface {
	GetAll(dto.QueryParam) []dto.ShowOrder
	GetByNumber(string) dto.ShowOrderDetail

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

func (o *OrderUsecase) GetAll(query dto.QueryParam) []dto.ShowOrder {

	gorm_db := o.OrderRepository.Filter()

	if query.Search != "" {
		gorm_db = o.OrderRepository.FilterSearch(gorm_db, query.Search)
	}

	if query.StatusID != "" {
		status_id, _ := strconv.ParseUint(query.StatusID, 10, 64)
		gorm_db = o.OrderRepository.FilterStatus(gorm_db, uint(status_id))
	}

	if query.CustomerID != "" {
		customer_id, _ := strconv.ParseUint(query.StatusID, 10, 64)
		gorm_db = o.OrderRepository.FilterCustomer(gorm_db, uint(customer_id))
	}

	if !query.From.IsZero() && !query.To.IsZero() {
		gorm_db = o.OrderRepository.FilterTime(gorm_db, query.From, query.To)
	}

	orders := o.OrderRepository.Exec(gorm_db)

	return mapper.ToShowOrderDtoList(orders)
}

func (o *OrderUsecase) GetByNumber(order_number string) dto.ShowOrderDetail {

	order := o.OrderRepository.GetByNumber(order_number)

	return mapper.ToShowOrderDetailDto(order)
}

func (o *OrderUsecase) Create(dto dto.Order) error {

	dto.Number = o.GenerateOrderNumber()
	dto.StatusID = 1

	order_entity := mapper.ToOrderEntity(dto)
	order_entity.Base = entity.BaseCreate()

	return o.OrderRepository.Create(order_entity)
}

func (o *OrderUsecase) Update(dto dto.Order) error {

	if !o.CheckOrderNumber(dto.Number) {
		return errors.New(config.OrderNotFound)
	}

	order_entity := mapper.ToOrderEntity(dto)
	order_entity.Base = entity.BaseUpdate()

	return o.OrderRepository.Update(order_entity)
}

func (o *OrderUsecase) GenerateOrderNumber() string {

	date := GetDate()
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)

	if len(date) == 5 {
		date = "0" + date
	}

	gorm_db := o.OrderRepository.Filter()
	gorm_db = o.OrderRepository.FilterTime(gorm_db, Bod(today), Bod(tomorrow))
	orders := o.OrderRepository.Exec(gorm_db)

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

func (o *OrderUsecase) CheckOrderNumber(order_number string) bool {

	order := o.OrderRepository.GetByNumber(order_number)

	if order.ID == 0 {
		return false
	}

	return true
}

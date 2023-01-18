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
	Create(entity.Order) (entity.Order, error)
	Update(entity.Order) error
}

type OrderUsecase struct {
	OrderRepository repository.OrderRepository
	ArticleUsecase  ArticleUsecase
}

func ProviderOrderUsecase(
	o repository.OrderRepository,
	a ArticleUsecase,
) OrderUsecase {
	return OrderUsecase{
		OrderRepository: o,
		ArticleUsecase:  a,
	}
}

func (o *OrderUsecase) Create(dto dto.Order) error {

	dto.Number = o.GenerateOrderNumber()

	order_entity := mapper.ToOrderEntity(dto)
	order_entity.Base = entity.BaseCreate()

	order, err := o.OrderRepository.Create(order_entity)
	if err != nil {
		return err
	}

	if err = o.ArticleUsecase.Check(order.ID, dto.Articles); err != nil {
		return err
	}

	return nil
}

func (o *OrderUsecase) GenerateOrderNumber() string {

	t := time.Now()
	return strconv.FormatInt(t.UnixNano(), 10)

}

package mapper

import (
	"order/dto"
	"order/entity"
)

func ToOrderEntity(dto dto.Order) entity.Order {
	return entity.Order{
		ID:           dto.ID,
		Number:       dto.Number,
		StatusID:     dto.StatusID,
		CustomerID:   dto.CustomerID,
		CustomerName: dto.CustomerName,
		Amount:       dto.Amount,
		Note:         dto.Note,
	}
}

func ToShowOrderDto(entity entity.Order) dto.ShowOrder {
	return dto.ShowOrder{
		ID:           entity.ID,
		Number:       entity.Number,
		CustomerName: entity.CustomerName,
		StatusName:   entity.Status.Name,
		Amount:       entity.Amount,
		ArticleCount: len(entity.Articles),
		Note:         entity.Note,
		OrderDate:    entity.Base.Created_At,
		OrderBy:      entity.Base.Created_By,
	}
}

func ToShowOrderDtoList(entity []entity.Order) []dto.ShowOrder {

	orders := make([]dto.ShowOrder, len(entity))

	for i, value := range entity {
		orders[i] = ToShowOrderDto(value)
	}

	return orders
}

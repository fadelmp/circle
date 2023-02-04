package mapper

import (
	"order/dto"
	"order/entity"
)

func ToOrderEntity(dto dto.Order) entity.Order {
	return entity.Order{
		ID:            dto.ID,
		Number:        dto.Number,
		StatusID:      dto.StatusID,
		CustomerID:    dto.CustomerID,
		CustomerName:  dto.CustomerName,
		Amount:        dto.Amount,
		Date:          dto.Date,
		Type:          dto.Type,
		DeliveryOrder: dto.DeliveryOrder,
		Note:          dto.Note,
		Articles:      ToArticleEntityList(dto.Articles),
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
		Date:         entity.Date,
		Type:         entity.Type,
		OrderBy:      entity.Base.Created_By,
	}
}

func ToShowOrderDetailDto(entity entity.Order) dto.ShowOrderDetail {
	return dto.ShowOrderDetail{
		ID:           entity.ID,
		Number:       entity.Number,
		CustomerID:   entity.CustomerID,
		CustomerName: entity.CustomerName,
		StatusName:   entity.Status.Name,
		Amount:       entity.Amount,
		ArticleCount: len(entity.Articles),
		Date:         entity.Date,
		Type:         entity.Type,
		OrderBy:      entity.Base.Created_By,
		Articles:     ToArticleDtoList(entity.Articles),
	}
}

func ToShowOrderDtoList(entity []entity.Order) []dto.ShowOrder {

	orders := make([]dto.ShowOrder, len(entity))

	for i, value := range entity {
		orders[i] = ToShowOrderDto(value)
	}

	return orders
}

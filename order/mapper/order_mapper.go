package mapper

import (
	"order/dto"
	"order/entity"
)

func ToOrderEntity(dto dto.Order) entity.Order {
	return entity.Order{
		ID:         dto.ID,
		Number:     dto.Number,
		StatusID:   dto.StatusID,
		CustomerID: dto.CustomerID,
		Amount:     dto.Amount,
		Note:       dto.Note,
	}
}

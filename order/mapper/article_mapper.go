package mapper

import (
	"order/dto"
	"order/entity"
)

func ToArticleEntity(dto dto.Article) entity.Article {
	return entity.Article{
		ID:         dto.ID,
		OrderID:    dto.OrderID,
		Name:       dto.Name,
		Quantity:   dto.Quantity,
		Unit:       dto.Unit,
		Amount:     dto.Amount,
		Note:       dto.Note,
		Path_Image: dto.Path_Image,
	}
}

package mapper

import (
	"order/dto"
	"order/entity"
)

func ToServiceEntity(dto dto.Service) entity.Service {
	return entity.Service{
		ID:        dto.ID,
		ArticleID: dto.ArticleID,
		ServiceID: dto.ServiceID,
		Amount:    dto.Amount,
	}
}

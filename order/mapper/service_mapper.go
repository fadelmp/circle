package mapper

import (
	"order/dto"
	"order/entity"
)

func ToServiceEntity(dto dto.Service) entity.Service {
	return entity.Service{
		ID:          dto.ID,
		ArticleID:   dto.ArticleID,
		ServiceID:   dto.ServiceID,
		ServiceName: dto.ServiceName,
		Amount:      dto.Amount,
	}
}

func ToServiceEntityList(dto []dto.Service) []entity.Service {

	services := make([]entity.Service, len(dto))

	for i, value := range dto {
		services[i] = ToServiceEntity(value)
	}

	return services

}

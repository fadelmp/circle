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

func ToServiceDto(entity entity.Service) dto.Service {
	return dto.Service{
		ID:          entity.ID,
		ServiceName: entity.ServiceName,
	}
}

func ToServiceEntityList(dto []dto.Service) []entity.Service {

	services := make([]entity.Service, len(dto))

	for i, value := range dto {
		services[i] = ToServiceEntity(value)
	}

	return services

}

func ToServiceDtoList(entity []entity.Service) []dto.Service {

	services := make([]dto.Service, len(entity))

	for i, value := range entity {
		services[i] = ToServiceDto(value)
	}

	return services
}

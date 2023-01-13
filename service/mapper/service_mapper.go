package mapper

import (
	"service/dto"
	"service/entity"
)

func ToServiceEntity(dto dto.Service) entity.Service {
	return entity.Service{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func ToServiceDto(entity entity.Service) dto.Service {
	return dto.Service{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
	}
}

func ToServiceDtoList(entity []entity.Service) []dto.Service {
	services := make([]dto.Service, len(entity))

	for i, value := range entity {
		services[i] = ToServiceDto(value)
	}

	return services
}

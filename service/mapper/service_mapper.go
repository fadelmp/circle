package mapper

import (
	"service/dto"
	"service/entity"
)

func ToServiceEntity(dto dto.Service, base entity.Base) entity.Service {
	return entity.Service{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Code:        dto.Code,
		Base:        base,
	}
}

func ToServiceDto(entity entity.Service) dto.Service {
	return dto.Service{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Code:        entity.Code,
		Price:       entity.Price,
		IsActived:   entity.Base.Is_Actived,
	}
}

func ToServiceDtoDetail(entity entity.Service) dto.ServiceDetail {
	return dto.ServiceDetail{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Code:        entity.Code,
		Price:       entity.Price,
		Base:        ToBaseDto(entity.Base),
	}
}

func ToServiceDtoList(entity []entity.Service) []dto.Service {
	services := make([]dto.Service, len(entity))

	for i, value := range entity {
		services[i] = ToServiceDto(value)
	}

	return services
}

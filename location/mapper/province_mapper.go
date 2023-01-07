package mapper

import (
	"location/dto"
	"location/entity"
)

func ToProvinceDto(entity entity.Province) dto.Province {
	return dto.Province{
		ID:        entity.ID,
		CountryID: entity.CountryID,
		Name:      entity.Name,
	}
}

func ToProvinceDtoList(entity []entity.Province) []dto.Province {
	provinces := make([]dto.Province, len(entity))

	for i, value := range entity {
		provinces[i] = ToProvinceDto(value)
	}

	return provinces
}

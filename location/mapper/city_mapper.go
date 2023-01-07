package mapper

import (
	"location/dto"
	"location/entity"
)

func ToCityDto(entity entity.City) dto.City {
	return dto.City{
		ID:           entity.ID,
		Name:         entity.Name,
		ProvinceID:   entity.ProvinceID,
		ProvinceName: entity.Province.Name,
	}
}

func ToCityDtoList(entity []entity.City) []dto.City {
	cities := make([]dto.City, len(entity))

	for i, value := range entity {
		cities[i] = ToCityDto(value)
	}

	return cities
}

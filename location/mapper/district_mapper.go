package mapper

import (
	"location/dto"
	"location/entity"
)

func ToDistrictDto(entity entity.District) dto.District {
	return dto.District{
		ID:       entity.ID,
		Name:     entity.Name,
		CityID:   entity.CityID,
		CityName: entity.City.Name,
	}
}

func ToDistrictDtoList(entity []entity.District) []dto.District {
	cities := make([]dto.District, len(entity))

	for i, value := range entity {
		cities[i] = ToDistrictDto(value)
	}

	return cities
}

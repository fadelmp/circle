package mapper

import (
	"location/dto"
	"location/entity"
)

func ToCountryDto(entity entity.Country) dto.Country {
	return dto.Country{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func ToCountryDtoList(entity []entity.Country) []dto.Country {
	countries := make([]dto.Country, len(entity))

	for i, value := range entity {
		countries[i] = ToCountryDto(value)
	}

	return countries
}

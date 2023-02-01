package mapper

import (
	"order/dto"
	"order/entity"
)

func ToUnitDto(entity entity.Unit) dto.Unit {
	return dto.Unit{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
	}
}

func ToUnitDtoList(entity []entity.Unit) []dto.Unit {
	Units := make([]dto.Unit, len(entity))

	for i, value := range entity {
		Units[i] = ToUnitDto(value)
	}

	return Units
}

package mapper

import (
	"order/dto"
	"order/entity"
)

func ToStatusDto(entity entity.Status) dto.Status {
	return dto.Status{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Total:       len(entity.Orders),
	}
}

func ToStatusDtoList(entity []entity.Status) []dto.Status {
	statuses := make([]dto.Status, len(entity))

	for i, value := range entity {
		statuses[i] = ToStatusDto(value)
	}

	return statuses
}

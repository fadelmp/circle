package mapper

import (
	"order/dto"
	"order/entity"
)

func ToUnitEntity(dto dto.Unit, base entity.Base) entity.Unit {
	return entity.Unit{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Base:        base,
	}
}

func ToUnitDto(entity entity.Unit) dto.Unit {
	return dto.Unit{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		IsActived:   entity.Base.Is_Actived,
	}
}

func ToUnitDtoList(entity []entity.Unit) []dto.Unit {
	Units := make([]dto.Unit, len(entity))

	for i, value := range entity {
		Units[i] = ToUnitDto(value)
	}

	return Units
}

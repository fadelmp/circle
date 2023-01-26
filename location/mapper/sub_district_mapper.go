package mapper

import (
	"location/dto"
	"location/entity"
)

func ToSubDistrictDto(entity entity.SubDistrict) dto.SubDistrict {
	return dto.SubDistrict{
		ID:           entity.ID,
		Name:         entity.Name,
		DistrictID:   entity.DistrictID,
		DistrictName: entity.District.Name,
	}
}

func ToSubDistrictDtoList(entity []entity.SubDistrict) []dto.SubDistrict {
	sub_districts := make([]dto.SubDistrict, len(entity))

	for i, value := range entity {
		sub_districts[i] = ToSubDistrictDto(value)
	}

	return sub_districts
}

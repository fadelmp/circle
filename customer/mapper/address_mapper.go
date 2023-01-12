package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToAddressEntity(dto dto.Address) entity.Address {
	return entity.Address{
		ID:         dto.ID,
		CustomerID: dto.CustomerID,
		Line:       dto.Line,
		DistrictID: dto.DistrictID,
		CityID:     dto.CityID,
		ProvinceID: dto.ProvinceID,
		CountryID:  dto.CountryID,
		PostalCode: dto.PostalCode,
	}
}

func ToAddressDto(entity entity.Address) dto.Address {
	return dto.Address{
		ID:         entity.ID,
		CustomerID: entity.CustomerID,
		Line:       entity.Line,
		DistrictID: entity.DistrictID,
		CityID:     entity.CityID,
		ProvinceID: entity.ProvinceID,
		CountryID:  entity.CountryID,
		PostalCode: entity.PostalCode,
	}
}

func ToAddressDtoList(entity []entity.Address) []dto.Address {
	addresses := make([]dto.Address, len(entity))

	for i, value := range entity {
		addresses[i] = ToAddressDto(value)
	}

	return addresses
}

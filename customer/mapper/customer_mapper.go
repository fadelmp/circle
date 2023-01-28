package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToCustomerEntity(dto dto.Customer, base entity.Base) entity.Customer {
	return entity.Customer{
		ID:         dto.ID,
		Name:       dto.Name,
		Phone:      dto.Phone,
		OtherPhone: dto.OtherPhone,
		Email:      dto.Email,
		Address:    Tes(dto.Address),
		Base:       base,
	}
}

func ToCustomerDto(entity entity.Customer, location string) dto.Customer {
	return dto.Customer{
		ID:          entity.ID,
		Name:        entity.Name,
		Phone:       entity.Phone,
		OtherPhone:  entity.OtherPhone,
		Email:       entity.Email,
		AddressLine: location,
		Address:     ToAddressDto(entity.Address),
		Base:        ToBaseDto(entity.Base),
	}
}

func ToCustomerDtoList(entity []entity.Customer, location []string) []dto.Customer {
	countries := make([]dto.Customer, len(entity))

	for i, value := range entity {
		countries[i] = ToCustomerDto(value, location[i])
	}

	return countries
}

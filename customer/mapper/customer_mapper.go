package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToCustomerEntity(dto dto.Customer) entity.Customer {
	return entity.Customer{
		ID:    dto.ID,
		Name:  dto.Name,
		Phone: dto.Phone,
		Email: dto.Email,
	}
}

func ToCustomerDto(entity entity.Customer) dto.Customer {
	return dto.Customer{
		ID:      entity.ID,
		Name:    entity.Name,
		Phone:   entity.Phone,
		Email:   entity.Email,
		Address: ToAddressDto(entity.Address),
		Company: ToCompanyDto(entity.Company),
		Base:    ToBaseDto(entity.Base),
	}
}

func ToCustomerDtoList(entity []entity.Customer) []dto.Customer {
	countries := make([]dto.Customer, len(entity))

	for i, value := range entity {
		countries[i] = ToCustomerDto(value)
	}

	return countries
}

package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToShowCustomerDto(entity entity.Customer, location string) dto.ShowCustomer {
	return dto.ShowCustomer{
		ID:           entity.ID,
		Name:         entity.Name,
		Phone:        entity.Phone,
		Email:        entity.Email,
		AddressLine:  location,
		CompanyName:  entity.Company.Name,
		CompanyPhone: entity.Company.Phone,
		IsActived:    entity.Base.Is_Actived,
	}
}

func ToShowCustomerDtoList(entity []entity.Customer, location []string) []dto.ShowCustomer {

	customers := make([]dto.ShowCustomer, len(entity))

	for i, value := range entity {
		customers[i] = ToShowCustomerDto(value, location[i])
	}

	return customers
}

package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToShowCustomerDto(entity entity.Customer, location string) dto.ShowCustomer {
	return dto.ShowCustomer{
		ID:          entity.ID,
		Name:        entity.Name,
		Phone:       entity.Phone,
		OtherPhone:  entity.OtherPhone,
		Email:       entity.Email,
		AddressLine: location,
		IsActived:   entity.Base.Is_Actived,
	}
}

func ToShowCustomerDtoList(entity []entity.Customer, location []string) []dto.ShowCustomer {

	customers := make([]dto.ShowCustomer, len(entity))

	for i, value := range entity {
		customers[i] = ToShowCustomerDto(value, location[i])
	}

	return customers
}

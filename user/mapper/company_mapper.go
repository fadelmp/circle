package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToCompanyEntity(dto dto.Company, customer_id uint) entity.Company {
	return entity.Company{
		ID:         dto.ID,
		CustomerID: customer_id,
		Name:       dto.Name,
		Phone:      dto.Phone,
	}
}

func ToCompanyDto(entity entity.Company) dto.Company {
	return dto.Company{
		ID:         entity.ID,
		CustomerID: entity.CustomerID,
		Name:       entity.Name,
		Phone:      entity.Phone,
	}
}

func ToCompanyDtoList(entity []entity.Company) []dto.Company {
	contact_peoples := make([]dto.Company, len(entity))

	for i, value := range entity {
		contact_peoples[i] = ToCompanyDto(value)
	}

	return contact_peoples
}

package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToContactPeopleEntity(dto dto.ContactPeople) entity.ContactPeople {
	return entity.ContactPeople{
		ID:         dto.ID,
		CustomerID: dto.CustomerID,
		Name:       dto.Name,
		Phone:      dto.Phone,
	}
}

func ToContactPeopleDto(entity entity.ContactPeople) dto.ContactPeople {
	return dto.ContactPeople{
		ID:         entity.ID,
		CustomerID: entity.CustomerID,
		Name:       entity.Name,
		Phone:      entity.Phone,
	}
}

func ToContactPeopleDtoList(entity []entity.ContactPeople) []dto.ContactPeople {
	contact_peoples := make([]dto.ContactPeople, len(entity))

	for i, value := range entity {
		contact_peoples[i] = ToContactPeopleDto(value)
	}

	return contact_peoples
}

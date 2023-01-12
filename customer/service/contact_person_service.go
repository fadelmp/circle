package service

import (
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
)

type ContactPeopleServiceContract interface {
	Create(entity.Address) error
}

type ContactPeopleService struct {
	ContactPeopleRepository repository.ContactPeopleRepository
}

func ProviderContactPeopleService(cp repository.ContactPeopleRepository) ContactPeopleService {
	return ContactPeopleService{ContactPeopleRepository: cp}
}

// Implementation

func (cp *ContactPeopleService) Create(dto dto.ContactPeople) error {

	// Map dto to entity
	contact_people_entity := mapper.ToContactPeopleEntity(dto)

	// Create address data
	_, err := cp.ContactPeopleRepository.Create(contact_people_entity)

	// Map entity to dto
	return err
}

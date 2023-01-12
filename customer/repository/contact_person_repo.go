package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type ContactPeopleRepositoryContract interface {
	Create(entity.ContactPeople) (entity.ContactPeople, error)
}

type ContactPeopleRepository struct {
	DB *gorm.DB
}

func ProviderContactPeopleRepository(DB *gorm.DB) ContactPeopleRepository {
	return ContactPeopleRepository{DB: DB}
}

// Implementation

func (cp *ContactPeopleRepository) Create(contact_people entity.ContactPeople) (entity.ContactPeople, error) {

	// Create ContactPeople
	err := cp.DB.Create(&contact_people).Error

	return contact_people, err
}

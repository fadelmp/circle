package repository

import (
	entity "location/entity"

	"github.com/jinzhu/gorm"
)

type CountryRepositoryContract interface {
	GetAll() []entity.Country
	GetByID(uint) entity.Country
}

type CountryRepository struct {
	DB *gorm.DB
}

func ProviderCountryRepository(DB *gorm.DB) CountryRepository {
	return CountryRepository{DB: DB}
}

// Implementation

func (c *CountryRepository) GetAll() []entity.Country {

	var countries []entity.Country

	// Find All Address
	c.DB.Find(&countries)

	return countries
}

func (c *CountryRepository) GetByID(id uint) entity.Country {

	var country entity.Country

	// Find Address By Id
	c.DB.Where("id=?", id).Find(&country)

	return country
}

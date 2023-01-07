package repository

import (
	entity "location/entity"

	"github.com/jinzhu/gorm"
)

type CityRepositoryContract interface {
	GetAll() []entity.City
	GetByID(uint) entity.City
	GetByProvinceID(uint) []entity.City
}

type CityRepository struct {
	DB *gorm.DB
}

func ProviderCityRepository(DB *gorm.DB) CityRepository {
	return CityRepository{DB: DB}
}

// Implementation

func (c *CityRepository) GetAll() []entity.City {

	var cities []entity.City
	var city entity.City

	// Find All Province
	c.DB.Model(&city).Preload("Province").Find(&cities)

	return cities
}

func (c *CityRepository) GetByID(id uint) entity.City {

	var city entity.City

	// Find City By Id
	c.DB.Model(&city).Preload("Province").Where("id=?", id).Find(&city)

	return city
}

func (c *CityRepository) GetByProvinceID(province_id uint) []entity.City {

	var cities []entity.City
	var city entity.City

	// Find City By Province Id
	c.DB.Model(&city).Preload("Province").Where("province_id=?", province_id).Find(&cities)

	return cities
}

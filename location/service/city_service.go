package service

import (
	"location/dto"
	"location/mapper"
	repository "location/repository"
)

type CityServiceContract interface {
	GetAll() []dto.City
	GetByID(uint) dto.City
	GetByprovinceID(uint) []dto.City
}

type CityService struct {
	CityRepository repository.CityRepository
}

func ProviderCityService(c repository.CityRepository) CityService {
	return CityService{CityRepository: c}
}

// Implementation

func (c *CityService) GetAll() []dto.City {

	// get all cities
	cities := c.CityRepository.GetAll()

	// map cities entity to cities dto
	return mapper.ToCityDtoList(cities)
}

func (c *CityService) GetByID(id uint) dto.City {

	// get city by id
	city := c.CityRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToCityDto(city)
}

func (c *CityService) GetByProvinceID(province_id uint) []dto.City {

	// get cities by id
	cities := c.CityRepository.GetByProvinceID(province_id)

	// map data from entity to dto
	return mapper.ToCityDtoList(cities)
}

package service

import (
	"location/dto"
	"location/mapper"
	repository "location/repository"
)

type CountryServiceContract interface {
	GetAll() []dto.Country
	GetByID(uint) dto.Country
}

type CountryService struct {
	CountryRepository repository.CountryRepository
}

func ProviderCountryService(c repository.CountryRepository) CountryService {
	return CountryService{CountryRepository: c}
}

// Implementation

func (c *CountryService) GetAll() []dto.Country {

	// get all address
	countries := c.CountryRepository.GetAll()

	// map address entity to address dto
	return mapper.ToCountryDtoList(countries)
}

func (c *CountryService) GetByID(id uint) dto.Country {

	// get address by id
	countries := c.CountryRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToCountryDto(countries)
}

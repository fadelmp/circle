package service

import (
	"location/dto"
	"location/mapper"
	repository "location/repository"
)

type DistrictServiceContract interface {
	GetAll() []dto.District
	GetByID(uint) dto.District
	GetByCityID(uint) []dto.District
}

type DistrictService struct {
	DistrictRepository repository.DistrictRepository
}

func ProviderDistrictService(c repository.DistrictRepository) DistrictService {
	return DistrictService{DistrictRepository: c}
}

// Implementation

func (d *DistrictService) GetAll() []dto.District {

	// get all cities
	districts := d.DistrictRepository.GetAll()

	// map cities entity to cities dto
	return mapper.ToDistrictDtoList(districts)
}

func (d *DistrictService) GetByID(id uint) dto.District {

	// get District by id
	district := d.DistrictRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToDistrictDto(district)
}

func (d *DistrictService) GetByCityID(city_id uint) []dto.District {

	// get cities by id
	districts := d.DistrictRepository.GetByCityID(city_id)

	// map data from entity to dto
	return mapper.ToDistrictDtoList(districts)
}

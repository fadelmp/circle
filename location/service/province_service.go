package service

import (
	"location/dto"
	"location/mapper"
	repository "location/repository"
)

type ProvinceServiceContract interface {
	GetAll() []dto.Province
	GetByID(uint) dto.Province
	GetByCountryID(uint) []dto.Province
}

type ProvinceService struct {
	ProvinceRepository repository.ProvinceRepository
}

func ProviderProvinceService(p repository.ProvinceRepository) ProvinceService {
	return ProvinceService{ProvinceRepository: p}
}

// Implementation

func (p *ProvinceService) GetAll() []dto.Province {

	// get all provinces
	provinces := p.ProvinceRepository.GetAll()

	// map provinces entity to provinces dto
	return mapper.ToProvinceDtoList(provinces)
}

func (p *ProvinceService) GetByID(id uint) dto.Province {

	// get province by id
	province := p.ProvinceRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToProvinceDto(province)
}

func (p *ProvinceService) GetByCountryID(country_id uint) []dto.Province {

	// get province by country id
	provinces := p.ProvinceRepository.GetByCountryID(country_id)

	// map data from entity to dto
	return mapper.ToProvinceDtoList(provinces)
}

package service

import (
	"location/dto"
	"location/mapper"
	repository "location/repository"
)

type SubDistrictServiceContract interface {
	GetAll() []dto.SubDistrict
	GetByID(uint) dto.SubDistrict
	GetByDistrictID(uint) []dto.SubDistrict
}

type SubDistrictService struct {
	SubDistrictRepository repository.SubDistrictRepository
}

func ProviderSubDistrictService(c repository.SubDistrictRepository) SubDistrictService {
	return SubDistrictService{SubDistrictRepository: c}
}

// Implementation

func (sd *SubDistrictService) GetAll() []dto.SubDistrict {

	// get all cities
	sub_districts := sd.SubDistrictRepository.GetAll()

	// map cities entity to cities dto
	return mapper.ToSubDistrictDtoList(sub_districts)
}

func (sd *SubDistrictService) GetByID(id uint) dto.SubDistrict {

	// get SubDistrict by id
	sub_district := sd.SubDistrictRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToSubDistrictDto(sub_district)
}

func (sd *SubDistrictService) GetByDistrictID(district_id uint) []dto.SubDistrict {

	// get cities by id
	sub_districts := sd.SubDistrictRepository.GetByDistrictID(district_id)

	// map data from entity to dto
	return mapper.ToSubDistrictDtoList(sub_districts)
}

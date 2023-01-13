package usecase

import (
	"errors"
	"service/config"
	"service/dto"
	entity "service/entity"
	"service/mapper"
	repository "service/repository"
)

type ServiceUsecaseContract interface {
	GetAll() []dto.Service
	GetByID(id uint) dto.Service

	Create(entity.Service) error
	Update(entity.Service) error
	Delete(id uint) error
	ActiveStatus(id uint) error
}

type ServiceUsecase struct {
	ServiceRepository repository.ServiceRepository
}

func ProviderServiceUsecase(s repository.ServiceRepository) ServiceUsecase {
	return ServiceUsecase{
		ServiceRepository: s,
	}
}

// Implementation

func (s *ServiceUsecase) Create(dto dto.Service) error {

	// check service name first, if name exists then return errors
	if !s.CheckName(dto) {
		return errors.New(config.ServiceExists)
	}

	// map dto to entity
	service_entity := mapper.ToServiceEntity(dto)

	// create service
	return s.ServiceRepository.Create(service_entity)
}

func (s *ServiceUsecase) CheckName(service dto.Service) bool {

	service_name := service.Name

	// get service by name
	service_data := s.ServiceRepository.GetByName(service_name)

	// return error if category exists
	if service_data.ID != 0 && service_data.Base.Is_Actived && !service_data.Base.Is_Deleted {
		return false
	}

	return true
}

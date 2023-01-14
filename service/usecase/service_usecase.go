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
	GetByID(uint) dto.Service

	Create(entity.Service) error
	Update(entity.Service) error
	Delete(uint) error
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

func (s *ServiceUsecase) GetAll() []dto.Service {

	services := s.ServiceRepository.GetAll()

	return mapper.ToServiceDtoList(services)
}

func (s *ServiceUsecase) GetByID(id uint) dto.Service {

	service := s.ServiceRepository.GetByID(id)

	return mapper.ToServiceDto(service)
}

func (s *ServiceUsecase) Create(dto dto.Service) error {

	// check service name first, if name exists then return errors
	if !s.CheckName(dto) {
		return errors.New(config.ServiceExists)
	}

	// map dto to entity
	service_entity := mapper.ToServiceEntity(dto)
	service_entity.Base = entity.BaseCreate()

	// create service
	return s.ServiceRepository.Create(service_entity)
}

func (s *ServiceUsecase) Update(dto dto.Service) error {

	if !s.CheckID(dto.ID) {
		return errors.New(config.ServiceNotFound)
	}

	service_entity := mapper.ToServiceEntity(dto)
	service_entity.Base = entity.BaseUpdate()

	return s.ServiceRepository.Update(service_entity)
}

func (s *ServiceUsecase) Delete(id uint) error {

	if !s.CheckID(id) {
		return errors.New(config.ServiceNotFound)
	}

	var service_entity entity.Service

	service_entity.ID = id
	service_entity.Base = entity.BaseDelete()

	return s.ServiceRepository.Delete(service_entity)
}

func (s *ServiceUsecase) CheckName(dto dto.Service) bool {

	service_name := dto.Name

	// get service by name
	service := s.ServiceRepository.GetByName(service_name)

	// return error if category exists
	if service.ID != 0 && service.Base.Is_Actived && !service.Base.Is_Deleted {
		return false
	}

	return true
}

func (s *ServiceUsecase) CheckID(id uint) bool {

	service := s.ServiceRepository.GetByID(id)

	if service.ID == 0 || service.Base.Is_Deleted {
		return false
	}

	return true
}

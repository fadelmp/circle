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
	GetAll(string, string) []dto.Service
	GetByID(uint) dto.Service

	Create(entity.Service) (error, int)
	Update(entity.Service) (error, int)
	Delete(uint) (error, int)
	Activate(uint, string) (error, int)
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

func (s *ServiceUsecase) GetAll(filter string, status string) []dto.Service {

	var services []entity.Service

	if filter != "" {
		services = s.ServiceRepository.GetByFilter(filter)
	} else if status == "available" {
		services = s.ServiceRepository.GetAvailable()
	} else if status == "active" {
		services = s.ServiceRepository.GetActive()
	} else {
		services = s.ServiceRepository.GetAll()
	}

	return mapper.ToServiceDtoList(services)
}

func (s *ServiceUsecase) GetByID(id uint) dto.Service {

	service := s.ServiceRepository.GetByID(id)

	return mapper.ToServiceDto(service)
}

func (s *ServiceUsecase) Create(dto dto.Service) (error, int) {

	// check service name first, if name exists then return errors
	if !s.CheckName(dto) {
		return errors.New(config.ServiceExists), 2
	}

	// map dto to entity
	service_entity := mapper.ToServiceEntity(dto)
	service_entity.Base = entity.BaseCreate()

	// create service
	return s.ServiceRepository.Create(service_entity), 0
}

func (s *ServiceUsecase) Update(dto dto.Service) (error, int) {

	if !s.CheckID(dto.ID) {
		return errors.New(config.ServiceNotFound), 1
	}

	if !s.CheckName(dto) {
		return errors.New(config.ServiceExists), 2
	}

	service_entity := mapper.ToServiceEntity(dto)
	service_entity.Base = entity.BaseUpdate()

	return s.ServiceRepository.Update(service_entity), 0
}

func (s *ServiceUsecase) Delete(id uint) (error, int) {

	if !s.CheckID(id) {
		return errors.New(config.ServiceNotFound), 1
	}

	var service_entity entity.Service
	service_entity.ID = id
	service_entity.Base = entity.BaseDelete()

	return s.ServiceRepository.ChangeStatus(service_entity), 0
}

func (s *ServiceUsecase) Activate(id uint, status string) (error, int) {

	if !s.CheckID(id) {
		return errors.New(config.ServiceNotFound), 1
	}

	is_active := true
	if status == "deactivate" {
		is_active = false
	}

	var service_entity entity.Service
	service_entity.ID = id
	service_entity.Base = entity.BaseActivate(is_active)

	return s.ServiceRepository.ChangeStatus(service_entity), 0
}

func (s *ServiceUsecase) CheckName(dto dto.Service) bool {

	service_name := dto.Name

	// get service by name
	service := s.ServiceRepository.GetByName(service_name)

	// return error if category exists
	if service.ID != 0 &&
		service.ID != dto.ID &&
		!service.Base.Is_Deleted {
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

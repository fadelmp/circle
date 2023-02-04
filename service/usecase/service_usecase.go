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
	GetAll(dto.QueryParam) []dto.Service
	GetByID(uint) dto.Service

	Create(entity.Service) (entity.Service, error, int)
	Update(entity.Service) (error, int)
	Delete(entity.Service) (error, int)
	Activate(entity.Service) (error, int)
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

func (s *ServiceUsecase) GetAll(dto dto.QueryParam) []dto.Service {

	var services []entity.Service

	if dto.Filter != "" {
		services = s.ServiceRepository.GetByFilter(dto.Filter)
	} else if dto.Status == "available" {
		services = s.ServiceRepository.GetAvailable()
	} else if dto.Status == "active" {
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

func (s *ServiceUsecase) Create(dto dto.Service) (dto.Service, error, int) {

	// check name whether name exists
	if !s.CheckName(dto) {
		return dto, errors.New(config.ServiceExists), 2
	}

	// map service dto to service entity
	service_entity := mapper.ToServiceEntity(dto, entity.BaseCreate())

	// create service and return
	service, err := s.ServiceRepository.Create(service_entity)

	return mapper.ToServiceDto(service), err, 0
}

func (s *ServiceUsecase) Update(dto dto.Service) (error, int) {

	// check id whether service not found
	if !s.CheckID(dto.ID) {
		return errors.New(config.ServiceNotFound), 1
	}

	// check name whether service name exists
	if !s.CheckName(dto) {
		return errors.New(config.ServiceExists), 2
	}

	// map service dto to service entity
	service_entity := mapper.ToServiceEntity(dto, entity.BaseUpdate())

	// update service and return
	return s.ServiceRepository.Update(service_entity), 0
}

func (s *ServiceUsecase) Delete(dto dto.Service) (error, int) {

	// check id whether service not exists
	if !s.CheckID(dto.ID) {
		return errors.New(config.ServiceNotFound), 1
	}

	// map service dto to service entity
	service_entity := mapper.ToServiceEntity(dto, entity.BaseDelete())

	// delete service and return
	return s.ServiceRepository.ChangeStatus(service_entity), 0
}

func (s *ServiceUsecase) Activate(dto dto.Service) (error, int) {

	// check id whether service not found
	if !s.CheckID(dto.ID) {
		return errors.New(config.ServiceNotFound), 1
	}

	// map service dto to service entity
	service_entity := mapper.ToServiceEntity(dto, entity.BaseActivate(dto.IsActived))

	// change service active status and return
	return s.ServiceRepository.ChangeStatus(service_entity), 0
}

func (s *ServiceUsecase) CheckName(dto dto.Service) bool {

	// get service by name
	service := s.ServiceRepository.GetByName(dto.Name)

	// return error if category exists
	if service.ID != 0 && service.ID != dto.ID &&
		!service.Base.Is_Deleted {
		return false
	}

	return true
}

func (s *ServiceUsecase) CheckID(id uint) bool {

	// get service by id
	service := s.ServiceRepository.GetByID(id)

	// if service not found return false
	if service.ID == 0 || service.Base.Is_Deleted {
		return false
	}

	return true
}

package usecase

import (
	"errors"
	"order/config"
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
)

type StatusUsecaseContract interface {
	GetAll() []dto.Status
	GetByID(id uint) dto.Status

	Create(entity.Status) error
	Update(entity.Status) error
	Delete(uint) error
	ActiveStatus(uint, bool) error
}

type StatusUsecase struct {
	StatusRepository repository.StatusRepository
}

func ProviderStatusUsecase(s repository.StatusRepository) StatusUsecase {
	return StatusUsecase{
		StatusRepository: s,
	}
}

func (s *StatusUsecase) GetAll() []dto.Status {

	statuses := s.StatusRepository.GetAll()

	return mapper.ToStatusDtoList(statuses)
}

func (s *StatusUsecase) GetByID(id uint) dto.Status {

	status := s.StatusRepository.GetByID(id)

	return mapper.ToStatusDto(status)
}

func (s *StatusUsecase) Create(dto dto.Status) error {

	if !s.CheckName(dto) {
		return errors.New(config.StatusExists)
	}

	status_entity := mapper.ToStatusEntity(dto)
	status_entity.Base = entity.BaseCreate()

	err := s.StatusRepository.Create(status_entity)

	return err
}

func (s *StatusUsecase) Update(dto dto.Status) error {

	if !s.CheckID(dto.ID) {
		return errors.New(config.StatusNotFound)
	}

	status_entity := mapper.ToStatusEntity(dto)
	status_entity.Base = entity.BaseUpdate()

	err := s.StatusRepository.Update(status_entity)

	return err
}

func (s *StatusUsecase) Delete(id uint) error {

	if !s.CheckID(id) {
		return errors.New(config.StatusNotFound)
	}

	var status_entity entity.Status

	status_entity.ID = id
	status_entity.Base = entity.BaseDelete()

	return s.StatusRepository.Delete(status_entity)
}

func (s *StatusUsecase) ActiveStatus(id uint, is_active bool) error {

	if !s.CheckID(id) {
		return errors.New(config.StatusNotFound)
	}

	var status_entity entity.Status

	status_entity.ID = id
	status_entity.Base = entity.BaseActivate(is_active)

	return s.StatusRepository.ActiveStatus(status_entity)
}

func (s *StatusUsecase) CheckName(dto dto.Status) bool {

	status_name := dto.Name

	status_data := s.StatusRepository.GetByName(status_name)

	if status_data.ID != 0 &&
		status_data.Base.Is_Actived &&
		!status_data.Base.Is_Deleted {
		return false
	}

	return true
}

func (s *StatusUsecase) CheckID(id uint) bool {

	status_data := s.StatusRepository.GetByID(id)

	if status_data.ID == 0 ||
		status_data.Base.Is_Deleted {
		return false
	}

	return true
}

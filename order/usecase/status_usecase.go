package usecase

import (
	"order/dto"
	"order/mapper"
	repository "order/repository"
)

type StatusUsecaseContract interface {
	GetAll() []dto.Status
	GetByID(id uint) dto.Status
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

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

	statuses_dto := mapper.ToStatusDtoList(statuses)

	return s.AddTotal(statuses_dto)
}

func (s *StatusUsecase) GetByID(id uint) dto.Status {

	status := s.StatusRepository.GetByID(id)

	return mapper.ToStatusDto(status)
}

func (s *StatusUsecase) AddTotal(statuses []dto.Status) []dto.Status {

	var total int
	for _, value := range statuses {
		total += value.Total
	}

	status := dto.Status{
		ID:    0,
		Name:  "Total",
		Total: total,
	}

	statuses = append(statuses, status)

	return statuses
}

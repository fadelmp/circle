package usecase

import (
	"order/dto"
	"order/mapper"
	repository "order/repository"
)

type ServiceUsecaseContract interface {
	Check(uint, []dto.Service) error
	Create(dto.Service) error
}

type ServiceUsecase struct {
	ServiceRepository repository.ServiceRepository
}

func ProviderServiceUsecase(s repository.ServiceRepository) ServiceUsecase {
	return ServiceUsecase{
		ServiceRepository: s,
	}
}

func (s *ServiceUsecase) Check(order_article_id uint, services []dto.Service) error {

	for _, value := range services {

		value.ArticleID = order_article_id

		if err := s.Create(value); err != nil {
			return err
		}
	}

	return nil
}

func (s *ServiceUsecase) Create(dto dto.Service) error {

	service_entity := mapper.ToServiceEntity(dto)

	err := s.ServiceRepository.Create(service_entity)

	return err
}

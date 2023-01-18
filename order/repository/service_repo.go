package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type ServiceRepositoryContract interface {
	Create(entity.Service) error
}

type ServiceRepository struct {
	DB *gorm.DB
}

func ProviderServiceRepository(DB *gorm.DB) ServiceRepository {
	return ServiceRepository{DB: DB}
}

func (s *ServiceRepository) Create(service entity.Service) error {

	err := s.DB.Create(&service).Error

	return err
}

package repository

import (
	entity "service/entity"

	"github.com/jinzhu/gorm"
)

type ServiceRepositoryContract interface {
	GetAll() []entity.Service
	GetByID(uint) entity.Service
	GetByName(string) entity.Service

	Create(entity.Service) error
	Update(entity.Service) error
	Delete(entity.Service) error
	ActiveStatus(entity.Service) error
}

type ServiceRepository struct {
	DB *gorm.DB
}

func ProviderServiceRepository(DB *gorm.DB) ServiceRepository {
	return ServiceRepository{DB: DB}
}

// Implementation

func (s *ServiceRepository) GetAll() []entity.Service {

	var services []entity.Service

	// Get Service All
	s.DB.Find(&services)

	return services
}

func (s *ServiceRepository) GetByID(id uint) entity.Service {

	var service entity.Service

	// Get Service By Id
	s.DB.Where("id=?", id).Find(&service)

	return service
}

func (s *ServiceRepository) GetByName(name string) entity.Service {

	var service entity.Service

	// Get Service by Name
	s.DB.Where("name=?", name).Find(&service)

	return service
}

func (s *ServiceRepository) Create(service entity.Service) error {

	// Create Service
	err := s.DB.Create(&service).Error

	return err
}

func (s *ServiceRepository) Update(service entity.Service) error {

	// update Service by id
	err := s.DB.Model(&service).Update(&service).Error

	return err
}

func (s *ServiceRepository) Delete(service entity.Service) error {

	// delete Service by id, by change is active value to false
	err := s.DB.Model(&service).Where("id=?", service.ID).Updates(map[string]interface{}{
		"is_actived": service.Base.Is_Actived,
		"is_deleted": service.Base.Is_Deleted,
		"updated_at": service.Base.Updated_At,
		"updated_by": service.Base.Updated_By,
	}).Error

	return err
}

func (s *ServiceRepository) ActiveStatus(service entity.Service) error {

	// delete Service by id, by change is active value to false
	err := s.DB.Model(&service).Where("id=?", service.ID).Updates(map[string]interface{}{
		"is_actived": service.Base.Is_Actived,
		"updated_at": service.Base.Updated_At,
		"updated_by": service.Base.Updated_By,
	}).Error

	return err
}

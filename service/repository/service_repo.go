package repository

import (
	"service/config"
	entity "service/entity"
	"strconv"

	"github.com/go-redis/redis"
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
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderServiceRepository(DB *gorm.DB, Redis *redis.Client) ServiceRepository {
	return ServiceRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (s *ServiceRepository) GetAll() []entity.Service {

	var services []entity.Service

	query := s.DB.Find(&services)
	keys := "services"

	// Get Service All
	config.CheckRedisQuery(s.Redis, query, keys)

	return services
}

func (s *ServiceRepository) GetByID(id uint) entity.Service {

	var service entity.Service

	query := s.DB.Where("id=?", id).Find(&service)
	keys := "service_id_" + strconv.FormatUint(uint64(id), 10)

	// Get Service By Id
	config.CheckRedisQuery(s.Redis, query, keys)

	return service
}

func (s *ServiceRepository) GetByName(name string) entity.Service {

	var service entity.Service

	query := s.DB.Where("name=?", name).Find(&service)
	keys := "service_name_" + name

	// Get Service by Name
	config.CheckRedisQuery(s.Redis, query, keys)

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

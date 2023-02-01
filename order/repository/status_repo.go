package repository

import (
	entity "order/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type StatusRepositoryContract interface {
	GetAll() []entity.Status
	GetByID(uint) entity.Status
}

type StatusRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderStatusRepository(DB *gorm.DB, Redis *redis.Client) StatusRepository {
	return StatusRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (s *StatusRepository) GetAll() []entity.Status {

	var statuses []entity.Status

	s.DB.Order("id asc").Find(&statuses)

	return statuses
}

func (s *StatusRepository) GetByID(id uint) entity.Status {

	var status entity.Status

	s.DB.Where("id=?", id).Find(&status)

	return status
}

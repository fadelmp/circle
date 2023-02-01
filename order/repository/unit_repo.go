package repository

import (
	entity "order/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type UnitRepositoryContract interface {
	GetAll() []entity.Unit
	GetByID(uint) entity.Unit
}

type UnitRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderUnitRepository(DB *gorm.DB, Redis *redis.Client) UnitRepository {
	return UnitRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (u *UnitRepository) GetAll() []entity.Unit {

	var units []entity.Unit

	u.DB.Order("id asc").Find(&units)

	return units
}

func (u *UnitRepository) GetByID(id uint) entity.Unit {

	var unit entity.Unit

	u.DB.Where("id=?", id).Find(&unit)

	return unit
}

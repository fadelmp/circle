package repository

import (
	"order/config"
	entity "order/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type UnitRepositoryContract interface {
	GetAll() []entity.Unit
	GetActive() []entity.Unit
	GetAvailable() []entity.Unit

	GetByID(uint) entity.Unit
	GetByName(string) entity.Unit

	Create(entity.Unit) error
	Update(entity.Unit) error
	ChangeStatus(entity.Unit) error
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

	query := u.DB.Order("id asc").Find(&units)
	keys := "units"

	// Get Unit All
	config.CheckRedisQuery(u.Redis, query, keys)

	return units
}

func (u *UnitRepository) GetActive() []entity.Unit {

	var units []entity.Unit

	query := u.DB.Order("id asc").Where("is_actived=?", true).Find(&units)
	keys := "units_active"

	config.CheckRedisQuery(u.Redis, query, keys)

	return units
}

func (u *UnitRepository) GetAvailable() []entity.Unit {

	var units []entity.Unit

	query := u.DB.Order("id asc").Where("is_deleted=?", false).Find(&units)
	keys := "units_available"

	config.CheckRedisQuery(u.Redis, query, keys)

	return units
}

func (u *UnitRepository) GetByID(id uint) entity.Unit {

	var unit entity.Unit

	query := u.DB.Where("id=?", id).Find(&unit)
	keys := "unit_id_" + strconv.FormatUint(uint64(id), 10)

	// Get Unit By Id
	config.CheckRedisQuery(u.Redis, query, keys)

	return unit
}

func (u *UnitRepository) GetByName(name string) entity.Unit {

	var unit entity.Unit

	query := u.DB.Where("name=?", name).Find(&unit)
	keys := "unit_name_" + name

	// Get Unit by Name
	config.CheckRedisQuery(u.Redis, query, keys)

	return unit
}

func (u *UnitRepository) Create(unit entity.Unit) error {

	// Create Unit
	return u.DB.Create(&unit).Error
}

func (u *UnitRepository) Update(unit entity.Unit) error {

	// update Unit by id
	return u.DB.Model(&unit).Update(&unit).Error
}

func (u *UnitRepository) ChangeStatus(unit entity.Unit) error {

	// change is actived and is deleted value
	return u.DB.Model(&unit).Where("id=?", unit.ID).Updates(map[string]interface{}{
		"is_actived": unit.Base.Is_Actived,
		"is_deleted": unit.Base.Is_Deleted,
		"updated_at": unit.Base.Updated_At,
		"updated_by": unit.Base.Updated_By,
	}).Error
}

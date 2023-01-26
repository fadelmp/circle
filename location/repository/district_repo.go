package repository

import (
	"location/config"
	entity "location/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type DistrictRepositoryContract interface {
	GetAll() []entity.District
	GetByID(uint) entity.District
	GetByDistrictID(uint) []entity.District
}

type DistrictRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderDistrictRepository(DB *gorm.DB, Redis *redis.Client) DistrictRepository {
	return DistrictRepository{DB: DB, Redis: Redis}
}

// Implementation

func (d *DistrictRepository) GetAll() []entity.District {

	var districts []entity.District

	query := d.DB.Model(&districts).Preload("Province").Find(&districts)
	keys := "districts"

	// Find All Province
	config.CheckRedisQuery(d.Redis, query, keys)

	return districts
}

func (d *DistrictRepository) GetByID(id uint) entity.District {

	var district entity.District
	query := d.DB.Model(&district).Preload("City").Where("id=?", id).Find(&district)
	keys := "district_" + strconv.FormatUint(uint64(id), 10)

	// Find District By Id
	config.CheckRedisQuery(d.Redis, query, keys)

	return district
}

func (d *DistrictRepository) GetByCityID(city_id uint) []entity.District {

	var districts []entity.District

	query := d.DB.Model(&districts).Preload("City").Where("city_id=?", city_id).Find(&districts)
	keys := "district_city_" + strconv.FormatUint(uint64(city_id), 10)

	// Find District By Id
	config.CheckRedisQuery(d.Redis, query, keys)

	return districts
}

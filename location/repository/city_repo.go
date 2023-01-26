package repository

import (
	"location/config"
	entity "location/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type CityRepositoryContract interface {
	GetAll() []entity.City
	GetByID(uint) entity.City
	GetByProvinceID(uint) []entity.City
}

type CityRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderCityRepository(DB *gorm.DB, Redis *redis.Client) CityRepository {
	return CityRepository{DB: DB, Redis: Redis}
}

// Implementation

func (c *CityRepository) GetAll() []entity.City {

	var cities []entity.City
	var city entity.City

	query := c.DB.Model(&city).Preload("Province").Find(&cities)
	keys := "cities"

	// Find All Province
	config.CheckRedisQuery(c.Redis, query, keys)

	return cities
}

func (c *CityRepository) GetByID(id uint) entity.City {

	var city entity.City
	query := c.DB.Model(&city).Preload("Province").Where("id=?", id).Find(&city)
	keys := "city_" + strconv.FormatUint(uint64(id), 10)

	// Find City By Id
	config.CheckRedisQuery(c.Redis, query, keys)

	return city
}

func (c *CityRepository) GetByProvinceID(province_id uint) []entity.City {

	var cities []entity.City
	var city entity.City

	query := c.DB.Model(&city).Preload("Province").Where("province_id=?", province_id).Find(&cities)
	keys := "city_province_" + strconv.FormatUint(uint64(province_id), 10)

	// Find City By Id
	config.CheckRedisQuery(c.Redis, query, keys)

	return cities
}

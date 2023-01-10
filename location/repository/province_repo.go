package repository

import (
	"location/config"
	entity "location/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type ProvinceRepositoryContract interface {
	GetAll() []entity.Province
	GetByID(uint) entity.Province
	GetByCountryID(uint) []entity.Province
}

type ProvinceRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderProvinceRepository(DB *gorm.DB, Redis *redis.Client) ProvinceRepository {
	return ProvinceRepository{DB: DB, Redis: Redis}
}

// Implementation

func (p *ProvinceRepository) GetAll() []entity.Province {

	var provinces []entity.Province
	var province entity.Province

	query := p.DB.Model(&province).Preload("Country").Find(&provinces)
	keys := "provinces"

	config.CheckRedisQuery(p.Redis, query, keys)

	return provinces
}

func (p *ProvinceRepository) GetByID(id uint) entity.Province {

	var province entity.Province

	query := p.DB.Model(&province).Preload("Country").Where("id=?", id).Find(&province)
	keys := "province_" + strconv.FormatUint(uint64(id), 10)

	config.CheckRedisQuery(p.Redis, query, keys)

	return province
}

func (p *ProvinceRepository) GetByCountryID(country_id uint) []entity.Province {

	var provinces []entity.Province
	var province entity.Province

	// Find Province By Country ID
	query := p.DB.Model(&province).Preload("Country").Where("country_id=?", country_id).Find(&provinces)
	keys := "province_country_" + strconv.FormatUint(uint64(country_id), 10)

	config.CheckRedisQuery(p.Redis, query, keys)

	return provinces
}

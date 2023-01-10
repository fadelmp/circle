package repository

import (
	"location/config"
	entity "location/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type CountryRepositoryContract interface {
	GetAll() []entity.Country
	GetByID(uint) entity.Country
}

type CountryRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderCountryRepository(DB *gorm.DB, Redis *redis.Client) CountryRepository {
	return CountryRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (c *CountryRepository) GetAll() []entity.Country {

	var countries []entity.Country
	query := c.DB.Find(&countries)
	keys := "countries"

	config.CheckRedisQuery(c.Redis, query, keys)

	return countries
}

func (c *CountryRepository) GetByID(id uint) entity.Country {

	var country entity.Country
	query := c.DB.Where("id=?", id).Find(&country)
	keys := "country_" + strconv.FormatUint(uint64(id), 10)

	// Find Country By Id
	config.CheckRedisQuery(c.Redis, query, keys)

	return country
}

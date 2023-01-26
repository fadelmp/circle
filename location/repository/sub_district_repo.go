package repository

import (
	"location/config"
	entity "location/entity"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type SubDistrictRepositoryContract interface {
	GetAll() []entity.SubDistrict
	GetByID(uint) entity.SubDistrict
	GetByDistrictID(uint) []entity.SubDistrict
}

type SubDistrictRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderSubDistrictRepository(DB *gorm.DB, Redis *redis.Client) SubDistrictRepository {
	return SubDistrictRepository{DB: DB, Redis: Redis}
}

// Implementation

func (sd *SubDistrictRepository) GetAll() []entity.SubDistrict {

	var sub_districts []entity.SubDistrict

	query := sd.DB.Model(&entity.SubDistrict{}).Preload("District").Find(&sub_districts)
	keys := "sub_districts"

	// Find All Province
	config.CheckRedisQuery(sd.Redis, query, keys)

	return sub_districts
}

func (sd *SubDistrictRepository) GetByID(id uint) entity.SubDistrict {

	var sub_district entity.SubDistrict
	query := sd.DB.Model(&sub_district).Preload("District").Where("id=?", id).Find(&sub_district)
	keys := "sub_district_" + strconv.FormatUint(uint64(id), 10)

	// Find District By Id
	config.CheckRedisQuery(sd.Redis, query, keys)

	return sub_district
}

func (sd *SubDistrictRepository) GetByDistrictID(district_id uint) []entity.SubDistrict {

	var sub_districts []entity.SubDistrict

	query := sd.DB.Model(&sub_districts).Preload("District").Where("district_id=?", district_id).Find(&sub_districts)
	keys := "sub_district_city_" + strconv.FormatUint(uint64(district_id), 10)

	// Find District By Id
	config.CheckRedisQuery(sd.Redis, query, keys)

	return sub_districts
}

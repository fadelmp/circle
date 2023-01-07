package repository

import (
	entity "location/entity"

	"github.com/jinzhu/gorm"
)

type ProvinceRepositoryContract interface {
	GetAll() []entity.Province
	GetByID(uint) entity.Province
}

type ProvinceRepository struct {
	DB *gorm.DB
}

func ProviderProvinceRepository(DB *gorm.DB) ProvinceRepository {
	return ProvinceRepository{DB: DB}
}

// Implementation

func (p *ProvinceRepository) GetAll() []entity.Province {

	var provinces []entity.Province
	var province entity.Province

	// Find All Province
	p.DB.Model(&province).Preload("Country").Find(&provinces)

	return provinces
}

func (p *ProvinceRepository) GetByID(id uint) entity.Province {

	var province entity.Province

	// Find Province By Id
	p.DB.Model(&province).Preload("Country").Find(&province)

	return province
}

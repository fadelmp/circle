package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type CompanyRepositoryContract interface {
	Create(entity.Company) error
	Update(entity.Company) error
}

type CompanyRepository struct {
	DB *gorm.DB
}

func ProviderCompanyRepository(DB *gorm.DB) CompanyRepository {
	return CompanyRepository{DB: DB}
}

// Implementation

func (c *CompanyRepository) Create(company entity.Company) error {

	// Create Company
	err := c.DB.Create(&company).Error

	return err
}

func (c *CompanyRepository) Update(company entity.Company) error {

	// Update Company
	err := c.DB.Model(&company).Update(&company).Error

	return err
}

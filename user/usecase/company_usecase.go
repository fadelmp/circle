package usecase

import (
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
)

type CompanyUsecaseContract interface {
	Create(entity.Company, uint) error
	Update(entity.Company, uint) error
}

type CompanyUsecase struct {
	CompanyRepository repository.CompanyRepository
}

func ProviderCompanyUsecase(c repository.CompanyRepository) CompanyUsecase {
	return CompanyUsecase{CompanyRepository: c}
}

// Implementation

func (c *CompanyUsecase) Create(dto dto.Company, customer_id uint) error {

	// Map dto to entity
	company_entity := mapper.ToCompanyEntity(dto, customer_id)

	// Create Company Data
	err := c.CompanyRepository.Create(company_entity)

	// Map entity to dto
	return err
}

func (c *CompanyUsecase) Update(dto dto.Company, customer_id uint) error {

	// Map dto to entity
	company_entity := mapper.ToCompanyEntity(dto, customer_id)

	// Update Company Data
	err := c.CompanyRepository.Update(company_entity)

	// Map entity to DTO
	return err
}

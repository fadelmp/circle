package usecase

import (
	"errors"
	"user/config"
	"user/dto"
	entity "user/entity"
	"user/mapper"
	repository "user/repository"
)

type PrivilegeUsecaseContract interface {
	GetAll() []dto.Privilege
	GetByID(uint) dto.Privilege

	Create(entity.Privilege) error
	Update(entity.Privilege) error
	Delete(entity.Privilege) error
	ActiveStatus(entity.Privilege) error
}

type PrivilegeUsecase struct {
	PrivilegeRepository repository.PrivilegeRepository
}

func ProviderPrivilegeUsecase(p repository.PrivilegeRepository) PrivilegeUsecase {
	return PrivilegeUsecase{
		PrivilegeRepository: p,
	}
}

// Implementation

func (p *PrivilegeUsecase) GetAll() []dto.Privilege {

	privileges := p.PrivilegeRepository.GetAll()

	return mapper.ToPrivilegeDtoList(privileges)
}

func (p *PrivilegeUsecase) GetByID(id uint) dto.Privilege {

	privilege := p.PrivilegeRepository.GetByID(id)

	return mapper.ToPrivilegeDto(privilege)
}

func (p *PrivilegeUsecase) Create(dto dto.Privilege) error {

	// change Privilege dto to entity to put on database
	privilege_entity := mapper.ToPrivilegeEntity(dto)
	privilege_entity.Base = entity.BaseCreate()

	// create Privilege data
	err := p.PrivilegeRepository.Create(privilege_entity)

	return err
}

func (c *PrivilegeUsecase) Update(dto dto.Privilege) error {

	if !c.CheckID(dto.ID) {
		return errors.New(config.PrivilegeNotFound)
	}

	Privilege_entity := mapper.ToPrivilegeEntity(dto)
	Privilege_entity.Base = entity.BaseUpdate()

	// create Privilege data
	err := c.PrivilegeRepository.Update(Privilege_entity)
	if err != nil {
		return err
	}

	// create address data
	if err := c.AddressUsecase.Update(dto.Address, dto.ID); err != nil {
		return err
	}

	// create contact person data
	if err := c.CompanyUsecase.Update(dto.Company, dto.ID); err != nil {
		return err
	}

	return nil
}

func (c *PrivilegeUsecase) Delete(id uint) error {

	if !c.CheckID(id) {
		return errors.New(config.PrivilegeNotFound)
	}

	var Privilege_entity entity.Privilege

	Privilege_entity.ID = id
	Privilege_entity.Base = entity.BaseDelete()

	return c.PrivilegeRepository.Delete(Privilege_entity)
}

func (c *PrivilegeUsecase) ActiveStatus(id uint, is_active bool) error {

	if !c.CheckID(id) {
		return errors.New(config.PrivilegeNotFound)
	}

	var Privilege_entity entity.Privilege

	Privilege_entity.ID = id
	Privilege_entity.Base = entity.BaseActivate(is_active)

	return c.PrivilegeRepository.ActiveStatus(Privilege_entity)
}

func (c *PrivilegeUsecase) CheckID(id uint) bool {

	Privilege_data := c.PrivilegeRepository.GetByID(id)

	if Privilege_data.ID == 0 ||
		!Privilege_data.Is_Actived ||
		Privilege_data.Is_Deleted {
		return false
	}

	return true
}

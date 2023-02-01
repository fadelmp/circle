package usecase

import (
	"errors"
	"order/config"
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
)

type UnitUsecaseContract interface {
	GetAll() []dto.Unit
	GetByID(uint) dto.Unit

	Create(entity.Unit) (error, int)
	Update(entity.Unit) (error, int)
	Delete(entity.Unit) (error, int)
	Activate(entity.Unit) (error, int)
}

type UnitUsecase struct {
	UnitRepository repository.UnitRepository
}

func ProviderUnitUsecase(s repository.UnitRepository) UnitUsecase {
	return UnitUsecase{
		UnitRepository: s,
	}
}

// Implementation

func (u *UnitUsecase) GetAll() []dto.Unit {

	var units []entity.Unit

	units = u.UnitRepository.GetAll()

	return mapper.ToUnitDtoList(units)
}

func (u *UnitUsecase) GetByID(id uint) dto.Unit {

	unit := u.UnitRepository.GetByID(id)

	return mapper.ToUnitDto(unit)
}

func (u *UnitUsecase) Create(dto dto.Unit) (error, int) {

	// check name whether name exists
	if !u.CheckName(dto) {
		return errors.New(config.UnitExists), 2
	}

	// map Unit dto to Unit entity
	unit_entity := mapper.ToUnitEntity(dto, entity.BaseCreate())

	// create Unit and return
	return u.UnitRepository.Create(unit_entity), 0
}

func (u *UnitUsecase) Update(dto dto.Unit) (error, int) {

	// check id whether Unit not found
	if !u.CheckID(dto.ID) {
		return errors.New(config.UnitNotFound), 1
	}

	// check name whether Unit name exists
	if !u.CheckName(dto) {
		return errors.New(config.UnitExists), 2
	}

	// map Unit dto to Unit entity
	unit_entity := mapper.ToUnitEntity(dto, entity.BaseUpdate())

	// update Unit and return
	return u.UnitRepository.Update(unit_entity), 0
}

func (u *UnitUsecase) Delete(dto dto.Unit) (error, int) {

	// check id whether Unit not exists
	if !u.CheckID(dto.ID) {
		return errors.New(config.UnitNotFound), 1
	}

	// map Unit dto to Unit entity
	unit_entity := mapper.ToUnitEntity(dto, entity.BaseDelete())

	// delete Unit and return
	return u.UnitRepository.ChangeStatus(unit_entity), 0
}

func (u *UnitUsecase) Activate(dto dto.Unit) (error, int) {

	// check id whether Unit not found
	if !u.CheckID(dto.ID) {
		return errors.New(config.UnitNotFound), 1
	}

	// map Unit dto to Unit entity
	unit_entity := mapper.ToUnitEntity(dto, entity.BaseActivate(dto.IsActived))

	// change Unit active status and return
	return u.UnitRepository.ChangeStatus(unit_entity), 0
}

func (u *UnitUsecase) CheckName(dto dto.Unit) bool {

	// get Unit by name
	unit := u.UnitRepository.GetByName(dto.Name)

	// return error if category exists
	if unit.ID != 0 && unit.ID != dto.ID &&
		!unit.Base.Is_Deleted {
		return false
	}

	return true
}

func (u *UnitUsecase) CheckID(id uint) bool {

	// get Unit by id
	unit := u.UnitRepository.GetByID(id)

	// if Unit not found return false
	if unit.ID == 0 || unit.Base.Is_Deleted {
		return false
	}

	return true
}

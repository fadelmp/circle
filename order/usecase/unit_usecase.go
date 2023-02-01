package usecase

import (
	"order/dto"
	"order/mapper"
	repository "order/repository"
)

type UnitUsecaseContract interface {
	GetAll() []dto.Unit
	GetByID(uint) dto.Unit
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

	units := u.UnitRepository.GetAll()

	return mapper.ToUnitDtoList(units)
}

func (u *UnitUsecase) GetByID(id uint) dto.Unit {

	unit := u.UnitRepository.GetByID(id)

	return mapper.ToUnitDto(unit)
}

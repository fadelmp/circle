package usecase

import (
	entity "customer/entity"
	"customer/request"
)

type LocationUsecaseContract interface {
	GetCountry(uint) string
	GetProvince(uint) string
	GetCity(uint) string
	GetDistrict(uint) string
	GetSubDistrict(uint64) string

	CheckLocation(entity.Customer) []string
}

type LocationUsecase struct {
	LocationRequest request.LocationRequest
}

func ProviderLocationUsecase(l request.LocationRequest) LocationUsecase {
	return LocationUsecase{
		LocationRequest: l,
	}
}

// Implementation

func (l *LocationUsecase) GetCountry(id uint) string {

	country := l.LocationRequest.GetCountry(id)

	return l.CheckValue(country)
}

func (l *LocationUsecase) GetProvince(id uint) string {

	province := l.LocationRequest.GetProvince(id)

	return l.CheckValue(province)
}

func (l *LocationUsecase) GetCity(id uint) string {

	city := l.LocationRequest.GetCity(id)

	return l.CheckValue(city)
}

func (l *LocationUsecase) GetDistrict(id uint) string {

	district := l.LocationRequest.GetDistrict(id)

	return l.CheckValue(district)
}

func (l *LocationUsecase) GetSubDistrict(id uint64) string {

	sub_district := l.LocationRequest.GetSubDistrict(id)

	return l.CheckValue(sub_district)
}

func (l *LocationUsecase) CheckValue(entity entity.Response) string {

	if entity.Result.Data.ID != 0 {
		return entity.Result.Data.Name
	}

	return ""
}

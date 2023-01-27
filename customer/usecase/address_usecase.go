package usecase

import (
	entity "customer/entity"
)

type AddressUsecaseContract interface {
	Check([]entity.Customer) []string
}

type AddressUsecase struct {
	LocationUsecase LocationUsecase
}

func ProviderAddressUsecase(l LocationUsecase) AddressUsecase {
	return AddressUsecase{
		LocationUsecase: l,
	}
}

// Implementation

func (a *AddressUsecase) Check(entity []entity.Customer) []string {

	var addresses []string

	for _, value := range entity {

		country := a.LocationUsecase.GetCountry(value.Address.CountryID)
		province := a.LocationUsecase.GetProvince(value.Address.ProvinceID)
		city := a.LocationUsecase.GetCity(value.Address.CityID)
		district := a.LocationUsecase.GetDistrict(value.Address.DistrictID)
		sub_district := a.LocationUsecase.GetSubDistrict(value.Address.SubDistrictID)

		address_line := value.Address.Line

		address_line += a.CheckComa(sub_district)
		address_line += a.CheckComa(district)
		address_line += a.CheckComa(city)
		address_line += a.CheckComa(province)
		address_line += a.CheckComa(country)
		address_line += a.CheckComa(value.Address.PostalCode)

		addresses = append(addresses, address_line)
	}

	return addresses
}

func (a *AddressUsecase) CheckComa(loc string) string {

	if loc != "" {
		loc = ", " + loc
	} else {
		loc = ""
	}

	return loc
}

package usecase

import (
	entity "customer/entity"
)

type AddressUsecaseContract interface {
	Check(entity.Customer) string
	CheckList([]entity.Customer) []string
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

func (a *AddressUsecase) Check(entity entity.Customer) string {

	country := a.LocationUsecase.GetCountry(entity.Address.CountryID)
	province := a.LocationUsecase.GetProvince(entity.Address.ProvinceID)
	city := a.LocationUsecase.GetCity(entity.Address.CityID)
	district := a.LocationUsecase.GetDistrict(entity.Address.DistrictID)
	sub_district := a.LocationUsecase.GetSubDistrict(entity.Address.SubDistrictID)

	address_line := entity.Address.Line

	address_line += a.CheckComa(sub_district, "Kel. ")
	address_line += a.CheckComa(district, "Kec. ")
	address_line += a.CheckComa(city, "")
	address_line += a.CheckComa(province, "")
	address_line += a.CheckComa(country, "")
	address_line += a.CheckComa(entity.Address.PostalCode, "")

	return address_line
}

func (a *AddressUsecase) CheckList(entity []entity.Customer) []string {

	var addresses []string

	for _, value := range entity {

		address_line := a.Check(value)
		addresses = append(addresses, address_line)
	}

	return addresses
}

func (a *AddressUsecase) CheckComa(loc string, header string) string {

	loc = ""

	if loc != "" {

		loc = ", " + loc
		if header != "" {
			loc = ", " + header + loc
		}
	}

	return loc
}

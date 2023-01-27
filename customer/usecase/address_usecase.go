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

		address_line := value.Address.Line + ", " +
			sub_district + ", " + district + ", " + city + ", " +
			province + ", " + country + ", " + value.Address.PostalCode
		addresses = append(addresses, address_line)
	}

	return addresses
}

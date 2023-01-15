package usecase

import (
	"core/dto"
	request "core/request"
	"os"
	"strconv"
)

type LocationUsecaseContract interface {
	GetAllCountry() dto.Response
	GetCountryByID(uint) dto.Response

	GetAllProvince() dto.Response
	GetProvinceByID(uint) dto.Response
	GetProvinceByCountryID(uint) dto.Response

	GetAllCity() dto.Response
	GetCityByID(uint) dto.Response
	GetCityByProvinceID(uint) dto.Response
}

type LocationUsecase struct {
	GetRequest request.GetRequest
}

func ProviderLocationUsecase(g request.GetRequest) LocationUsecase {
	return LocationUsecase{GetRequest: g}
}

func getUri() string {
	return os.Getenv("LOCATION_Usecase_URI")
}

// Implementation

func (l *LocationUsecase) GetAllCountry() dto.Response {

	uri := getUri()
	uri += "/country"

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetCountryByID(id uint) dto.Response {

	uri := getUri()
	uri += "/country/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetAllProvince() dto.Response {

	uri := getUri()
	uri += "/province"

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetProvinceByID(id uint) dto.Response {

	uri := getUri()
	uri += "/province/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetProvinceByCountryID(country_id uint) dto.Response {

	uri := getUri()
	uri += "/province/country/"
	uri += strconv.FormatUint(uint64(country_id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetAllCity() dto.Response {

	uri := getUri()
	uri += "/city"

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetCityByID(id uint) dto.Response {

	uri := getUri()
	uri += "/city/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationUsecase) GetCityByProvinceID(province_id uint) dto.Response {

	uri := getUri()
	uri += "/city/province/"
	uri += strconv.FormatUint(uint64(province_id), 10)

	return l.GetRequest.Main(uri)
}

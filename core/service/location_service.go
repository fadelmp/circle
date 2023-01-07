package service

import (
	"core/dto"
	request "core/request"
	"os"
	"strconv"
)

type LocationServiceContract interface {
	GetAllCountry() dto.Response
	GetCountryByID(uint) dto.Response

	GetAllProvince() dto.Response
	GetProvinceByID(uint) dto.Response
	GetProvinceByCountryID(uint) dto.Response

	GetAllCity() dto.Response
	GetCityByID(uint) dto.Response
	GetCityByProvinceID(uint) dto.Response
}

type LocationService struct {
	GetRequest request.GetRequest
}

func ProviderLocationService(g request.GetRequest) LocationService {
	return LocationService{GetRequest: g}
}

// Implementation

func (l *LocationService) GetAllCountry() dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/country"

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetCountryByID(id uint) dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/country/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetAllProvince() dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/province"

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetProvinceByID(id uint) dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/province/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetProvinceByCountryID(country_id uint) dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/province/country/"
	uri += strconv.FormatUint(uint64(country_id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetAllCity() dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/city"

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetCityByID(id uint) dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/city/"
	uri += strconv.FormatUint(uint64(id), 10)

	return l.GetRequest.Main(uri)
}

func (l *LocationService) GetCityByProvinceID(province_id uint) dto.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/city/province/"
	uri += strconv.FormatUint(uint64(province_id), 10)

	return l.GetRequest.Main(uri)
}

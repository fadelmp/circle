package request

import (
	"customer/entity"
	"os"
	"strconv"
)

type LocationRequestContract interface {
	GetCountry(uint) string
	GetProvince(uint) string
	GetCity(uint) string
	GetDistrict(uint) string
}

type LocationRequest struct {
	HTTPRequest HTTPRequest
}

func ProviderLocationRequest() LocationRequest { return LocationRequest{} }

func (l *LocationRequest) GetCountry(country_id uint) string {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/country/"
	uri += strconv.FormatUint(uint64(country_id), 10)

	response := l.HTTPRequest.Get(uri)
	return l.CheckValue(response)
}

func (l *LocationRequest) GetProvince(province_id uint) string {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/province/"
	uri += strconv.FormatUint(uint64(province_id), 10)

	response := l.HTTPRequest.Get(uri)
	return l.CheckValue(response)
}

func (l *LocationRequest) GetCity(city_id uint) string {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/city/"
	uri += strconv.FormatUint(uint64(city_id), 10)

	response := l.HTTPRequest.Get(uri)
	return l.CheckValue(response)
}

func (l *LocationRequest) GetDistrict(district_id uint) string {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/district/"
	uri += strconv.FormatUint(uint64(district_id), 10)

	response := l.HTTPRequest.Get(uri)
	return l.CheckValue(response)
}

func (l *LocationRequest) CheckValue(entity entity.Response) string {

	if entity.Result.Data.ID != 0 {
		return entity.Result.Data.Name
	}

	return ""
}

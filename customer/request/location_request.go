package request

import (
	"customer/entity"
	"os"
	"strconv"
)

type LocationRequestContract interface {
	GetCountry(uint) entity.Response
	GetProvince(uint) entity.Response
	GetCity(uint) entity.Response
	GetDistrict(uint) entity.Response
	GetSubDistrict(uint) entity.Response
}

type LocationRequest struct {
	HTTPRequest HTTPRequest
}

func ProviderLocationRequest() LocationRequest { return LocationRequest{} }

func (l *LocationRequest) GetCountry(country_id uint) entity.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/country/"
	uri += strconv.FormatUint(uint64(country_id), 10)

	return l.HTTPRequest.Get(uri)
}

func (l *LocationRequest) GetProvince(province_id uint) entity.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/province/"
	uri += strconv.FormatUint(uint64(province_id), 10)

	return l.HTTPRequest.Get(uri)
}

func (l *LocationRequest) GetCity(city_id uint) entity.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/city/"
	uri += strconv.FormatUint(uint64(city_id), 10)

	return l.HTTPRequest.Get(uri)
}

func (l *LocationRequest) GetDistrict(district_id uint) entity.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/district/"
	uri += strconv.FormatUint(uint64(district_id), 10)

	return l.HTTPRequest.Get(uri)
}

func (l *LocationRequest) GetSubDistrict(sub_district_id uint) entity.Response {

	uri := os.Getenv("LOCATION_SERVICE_URI")
	uri += "/sub_district/"
	uri += strconv.FormatUint(uint64(sub_district_id), 10)

	return l.HTTPRequest.Get(uri)
}

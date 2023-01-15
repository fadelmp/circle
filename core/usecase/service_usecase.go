package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strconv"
)

type ServiceUsecaseContract interface {
	GetServices() dto.Response
	GetServiceById(uint) dto.Response

	CreateService(interface{}) dto.Response
	UpdateService(interface{}) dto.Response
	DeleteService(uint) dto.Response
	ActiveStatus(uint, bool) dto.Response
}

type ServiceUsecase struct {
	GetRequest    request.GetRequest
	PutRequest    request.PutRequest
	PostRequest   request.PostRequest
	PatchRequest  request.PatchRequest
	DeleteRequest request.DeleteRequest
}

func ProviderServiceUsecase(
	g request.GetRequest,
	p request.PutRequest,
	po request.PostRequest,
	pa request.PatchRequest,
	d request.DeleteRequest,
) ServiceUsecase {
	return ServiceUsecase{
		GetRequest:    g,
		PutRequest:    p,
		PostRequest:   po,
		PatchRequest:  pa,
		DeleteRequest: d,
	}
}

func getServiceUri() string {

	uri := os.Getenv("SERVICE_URI")
	uri += "/service"

	return uri
}

// Implementation

func (s *ServiceUsecase) GetServices() dto.Response {

	uri := getServiceUri()
	return s.GetRequest.Main(uri)
}

func (s *ServiceUsecase) GetServiceById(id uint) dto.Response {

	uri := getServiceUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return s.GetRequest.Main(uri)
}

func (s *ServiceUsecase) CreateService(form_data interface{}) dto.Response {

	uri := getServiceUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return s.PostRequest.Main(uri, request_body)
}

func (s *ServiceUsecase) UpdateService(form_data interface{}) dto.Response {

	uri := getServiceUri()

	put_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(put_body)

	return s.PostRequest.Main(uri, request_body)
}

func (s *ServiceUsecase) DeleteService(id uint) dto.Response {

	uri := getServiceUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	var req_body *bytes.Buffer
	return s.DeleteRequest.Main(uri, req_body)
}
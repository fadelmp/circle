package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strconv"
)

type OrderUnitUsecaseContract interface {
	GetOrderUnits(string, string) dto.Response
	GetOrderUnitById(uint) dto.Response

	CreateOrderUnit(interface{}) dto.Response
	UpdateOrderUnit(interface{}) dto.Response
	DeleteOrderUnit(uint) dto.Response

	ActivateOrderUnit(uint, string) dto.Response
}

type OrderUnitUsecase struct {
	GetRequest    request.GetRequest
	PutRequest    request.PutRequest
	PostRequest   request.PostRequest
	PatchRequest  request.PatchRequest
	DeleteRequest request.DeleteRequest
}

func ProviderOrderUnitUsecase(
	g request.GetRequest,
	p request.PutRequest,
	po request.PostRequest,
	pa request.PatchRequest,
	d request.DeleteRequest,
) OrderUnitUsecase {
	return OrderUnitUsecase{
		GetRequest:    g,
		PutRequest:    p,
		PostRequest:   po,
		PatchRequest:  pa,
		DeleteRequest: d,
	}
}

func getOrderUnitUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order_unit"

	return uri
}

// Implementation

func (s *OrderUnitUsecase) GetOrderUnits(filter string, status string) dto.Response {

	uri := getOrderUnitUri()

	if filter != "" {
		uri += "?filter=" + filter
	} else if status != "" {
		uri += "?status=" + status
	}

	return s.GetRequest.Main(uri)
}

func (s *OrderUnitUsecase) GetOrderUnitById(id uint) dto.Response {

	uri := getOrderUnitUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return s.GetRequest.Main(uri)
}

func (s *OrderUnitUsecase) CreateOrderUnit(form_data interface{}) dto.Response {

	uri := getOrderUnitUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return s.PostRequest.Main(uri, request_body)
}

func (s *OrderUnitUsecase) UpdateOrderUnit(form_data interface{}) dto.Response {

	uri := getOrderUnitUri()

	put_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(put_body)

	return s.PutRequest.Main(uri, request_body)
}

func (s *OrderUnitUsecase) DeleteOrderUnit(id uint) dto.Response {

	uri := getOrderUnitUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return s.DeleteRequest.Main(uri)
}

func (s *OrderUnitUsecase) ActivateOrderUnit(id uint, status string) dto.Response {

	uri := getOrderUnitUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10) + "/" + status

	return s.PatchRequest.Main(uri)
}

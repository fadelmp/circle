package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strconv"
)

type OrderUsecaseContract interface {
	GetOrders() dto.Response
	GetOrderByOrderNumber(string) dto.Response
	GetOrderByCustomer(uint) dto.Response
	GetOrderByStatus(uint) dto.Response

	CreateOrder(interface{}) dto.Response
	UpdateOrder(interface{}) dto.Response
}

type OrderUsecase struct {
	GetRequest   request.GetRequest
	PutRequest   request.PutRequest
	PostRequest  request.PostRequest
	PatchRequest request.PatchRequest
}

func ProviderOrderUsecase(
	g request.GetRequest,
	pu request.PutRequest,
	po request.PostRequest,
	pa request.PatchRequest,
) OrderUsecase {
	return OrderUsecase{
		GetRequest:   g,
		PutRequest:   pu,
		PostRequest:  po,
		PatchRequest: pa,
	}
}

func getOrderUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order"

	return uri
}

// Implementation

func (o *OrderUsecase) GetOrders() dto.Response {

	uri := getOrderUri()

	return o.GetRequest.Main(uri)
}

func (o *OrderUsecase) GetOrderByOrderNumber(number string) dto.Response {

	uri := getOrderUri()
	uri += "/number/" + number

	return o.GetRequest.Main(uri)
}

func (o *OrderUsecase) GetOrderByCustomer(customer_id uint) dto.Response {

	uri := getOrderUri()
	uri += "/customer/" + strconv.FormatUint(uint64(customer_id), 10)

	return o.GetRequest.Main(uri)
}

func (o *OrderUsecase) GetOrderByStatus(status_id uint) dto.Response {

	uri := getOrderUri()
	uri += "/status/" + strconv.FormatUint(uint64(status_id), 10)

	return o.GetRequest.Main(uri)
}

func (o *OrderUsecase) CreateOrder(form_data interface{}) dto.Response {

	uri := getOrderUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return o.PostRequest.Main(uri, request_body)
}

func (o *OrderUsecase) UpdateOrder(form_data interface{}) dto.Response {

	uri := getOrderUri()

	put_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(put_body)

	return o.PutRequest.Main(uri, request_body)
}

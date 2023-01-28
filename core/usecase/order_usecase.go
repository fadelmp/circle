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
	GetOrderByCustomer(uint) dto.Response

	CreateOrder(interface{}) dto.Response
}

type OrderUsecase struct {
	GetRequest   request.GetRequest
	PostRequest  request.PostRequest
	PatchRequest request.PatchRequest
}

func ProviderOrderUsecase(
	g request.GetRequest,
	po request.PostRequest,
	pa request.PatchRequest,
) OrderUsecase {
	return OrderUsecase{
		GetRequest:   g,
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

func (os *OrderUsecase) GetOrders() dto.Response {

	uri := getOrderUri()

	return os.GetRequest.Main(uri)
}

func (os *OrderUsecase) GetOrderByCustomer(customer_id uint) dto.Response {

	uri := getOrderUri()
	uri += "/" + strconv.FormatUint(uint64(customer_id), 10)

	return os.GetRequest.Main(uri)
}

func (os *OrderUsecase) CreateOrder(form_data interface{}) dto.Response {

	uri := getOrderUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return os.PostRequest.Main(uri, request_body)
}

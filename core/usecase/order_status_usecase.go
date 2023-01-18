package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strconv"
)

type OrderStatusUsecaseContract interface {
	GetOrderStatuses() dto.Response
	GetOrderStatusesById(uint) dto.Response

	CreateOrderStatus(interface{}) dto.Response
	UpdateOrderStatus(interface{}) dto.Response
	DeleteOrderStatus(uint) dto.Response

	ActivateOrderStatus(uint) dto.Response
	DeactivateOrderStatus(uint) dto.Response
}

type OrderStatusUsecase struct {
	GetRequest    request.GetRequest
	PutRequest    request.PutRequest
	PostRequest   request.PostRequest
	PatchRequest  request.PatchRequest
	DeleteRequest request.DeleteRequest
}

func ProviderOrderStatusUsecase(
	g request.GetRequest,
	p request.PutRequest,
	po request.PostRequest,
	pa request.PatchRequest,
	d request.DeleteRequest,
) OrderStatusUsecase {
	return OrderStatusUsecase{
		GetRequest:    g,
		PutRequest:    p,
		PostRequest:   po,
		PatchRequest:  pa,
		DeleteRequest: d,
	}
}

func getOrderStatusUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order"

	return uri
}

// Implementation

func (os *OrderStatusUsecase) GetOrderStatuses() dto.Response {

	uri := getOrderStatusUri()
	return os.GetRequest.Main(uri)
}

func (os *OrderStatusUsecase) GetOrderStatusById(id uint) dto.Response {

	uri := getOrderStatusUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return os.GetRequest.Main(uri)
}

func (os *OrderStatusUsecase) CreateOrderStatus(form_data interface{}) dto.Response {

	uri := getOrderStatusUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return os.PostRequest.Main(uri, request_body)
}

func (os *OrderStatusUsecase) UpdateOrderStatus(form_data interface{}) dto.Response {

	uri := getOrderStatusUri()

	put_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(put_body)

	return os.PostRequest.Main(uri, request_body)
}

func (os *OrderStatusUsecase) DeleteOrderStatus(id uint) dto.Response {

	uri := getOrderStatusUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return os.DeleteRequest.Main(uri)
}

func (os *OrderStatusUsecase) ActivateOrderStatus(id uint) dto.Response {

	uri := getOrderStatusUri()
	uri += "/activate/" + strconv.FormatUint(uint64(id), 10)

	return os.PatchRequest.Main(uri)
}

func (os *OrderStatusUsecase) DeactivateOrderStatus(id uint) dto.Response {

	uri := getOrderStatusUri()
	uri += "/deactivate/" + strconv.FormatUint(uint64(id), 10)

	return os.PatchRequest.Main(uri)
}

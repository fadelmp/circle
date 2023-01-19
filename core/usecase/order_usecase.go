package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
)

type OrderUsecaseContract interface {
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

func (os *OrderUsecase) CreateOrder(form_data interface{}) dto.Response {

	uri := getOrderUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return os.PostRequest.Main(uri, request_body)
}

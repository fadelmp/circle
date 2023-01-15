package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
)

type CustomerUsecaseContract interface {
	CreateCustomer() dto.Response
}

type CustomerUsecase struct {
	GetRequest  request.GetRequest
	PostRequest request.PostRequest
}

func ProviderCustomerUsecase(
	g request.GetRequest,
	p request.PostRequest,
) CustomerUsecase {
	return CustomerUsecase{
		GetRequest:  g,
		PostRequest: p,
	}
}

// Implementation

func (c *CustomerUsecase) CreateCustomer(form_data interface{}) dto.Response {

	uri := os.Getenv("CUSTOMER_Usecase_URI")
	uri += "/customer"

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return c.PostRequest.Main(uri, request_body)
}

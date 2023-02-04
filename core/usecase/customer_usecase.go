package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type CustomerUsecaseContract interface {
	GetCustomers(string, string, string) dto.Response
	GetCustomerById(uint) dto.Response

	CreateCustomer(interface{}) dto.Response
	UpdateCustomer(interface{}) dto.Response
	DeleteCustomer(uint) dto.Response

	ActivateCustomer(uint, string) dto.Response
}

type CustomerUsecase struct {
	GetRequest    request.GetRequest
	PutRequest    request.PutRequest
	PostRequest   request.PostRequest
	PatchRequest  request.PatchRequest
	DeleteRequest request.DeleteRequest
}

func ProviderCustomerUsecase(
	g request.GetRequest,
	p request.PutRequest,
	po request.PostRequest,
	pa request.PatchRequest,
	d request.DeleteRequest,
) CustomerUsecase {
	return CustomerUsecase{
		GetRequest:    g,
		PutRequest:    p,
		PostRequest:   po,
		PatchRequest:  pa,
		DeleteRequest: d,
	}
}

func getCustomerUri() string {

	uri := os.Getenv("CUSTOMER_URI")
	uri += "/customer"

	return uri
}

// Implementation

func (c *CustomerUsecase) GetCustomers(filter string, status string, name string) dto.Response {

	uri := getCustomerUri()

	if filter != "" {
		uri += "?filter=" + filter
	} else if name != "" {
		uri += "?name=" + name
	} else if status != "" {
		uri += "?status=" + status
	}

	new_uri := strings.Replace(uri, " ", "%20", 3)

	return c.GetRequest.Main(new_uri)
}

func (c *CustomerUsecase) GetCustomerById(id uint) dto.Response {

	uri := getCustomerUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return c.GetRequest.Main(uri)
}

func (c *CustomerUsecase) CreateCustomer(form_data interface{}) dto.Response {

	uri := getCustomerUri()

	post_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(post_body)

	return c.PostRequest.Main(uri, request_body)
}

func (c *CustomerUsecase) UpdateCustomer(form_data interface{}) dto.Response {

	uri := getCustomerUri()

	put_body, _ := json.Marshal(form_data)
	request_body := bytes.NewBuffer(put_body)

	return c.PutRequest.Main(uri, request_body)
}

func (c *CustomerUsecase) DeleteCustomer(id uint) dto.Response {

	uri := getCustomerUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return c.DeleteRequest.Main(uri)
}

func (c *CustomerUsecase) ActivateCustomer(id uint, status string) dto.Response {

	uri := getCustomerUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10) + "/" + status

	return c.PatchRequest.Main(uri)
}

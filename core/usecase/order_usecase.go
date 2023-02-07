package usecase

import (
	"bytes"
	"core/dto"
	request "core/request"
	"encoding/json"
	"os"
	"strings"
)

type OrderUsecaseContract interface {
	GetOrders(dto.QueryParam) dto.Response
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

func (o *OrderUsecase) GetOrders(query_param dto.QueryParam) dto.Response {

	uri := getOrderUri()

	if query_param.Search != "" {
		uri = o.AddQuestionMark(uri)
		uri = uri + "search=" + query_param.Search
	}

	if query_param.StatusID != "" {
		uri = o.AddQuestionMark(uri)
		uri = uri + "status_id=" + query_param.StatusID
	}

	if query_param.CustomerID != "" {
		uri = o.AddQuestionMark(uri)
		uri = uri + "customer_id=" + query_param.CustomerID
	}

	if !query_param.From.IsZero() && !query_param.To.IsZero() {
		uri = o.AddQuestionMark(uri)
		uri = uri + "from=" + query_param.From.String()
		uri = uri + "&to=" + query_param.To.String()
	}

	return o.GetRequest.Main(uri)
}

func (o *OrderUsecase) GetOrderByOrderNumber(number string) dto.Response {

	uri := getOrderUri()
	uri += "/number/" + number

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

func (o *OrderUsecase) AddQuestionMark(uri string) string {

	if !strings.Contains(uri, "?") {
		uri = uri + "?"
	} else {
		uri = "&" + uri
	}

	return uri
}

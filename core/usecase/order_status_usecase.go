package usecase

import (
	"core/dto"
	request "core/request"
	"os"
	"strconv"
)

type OrderStatusUsecaseContract interface {
	GetOrderStatuses() dto.Response
	GetOrderStatusesById(uint) dto.Response
}

type OrderStatusUsecase struct {
	GetRequest request.GetRequest
}

func ProviderOrderStatusUsecase(g request.GetRequest) OrderStatusUsecase {
	return OrderStatusUsecase{
		GetRequest: g,
	}
}

func getOrderStatusUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order_status"

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

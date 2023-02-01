package usecase

import (
	"core/dto"
	request "core/request"
	"os"
	"strconv"
)

type OrderUnitUsecaseContract interface {
	GetOrderUnits(string, string) dto.Response
	GetOrderUnitById(uint) dto.Response
}

type OrderUnitUsecase struct {
	GetRequest request.GetRequest
}

func ProviderOrderUnitUsecase(g request.GetRequest) OrderUnitUsecase {
	return OrderUnitUsecase{
		GetRequest: g,
	}
}

func getOrderUnitUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order_unit"

	return uri
}

// Implementation

func (s *OrderUnitUsecase) GetOrderUnits() dto.Response {

	uri := getOrderUnitUri()

	return s.GetRequest.Main(uri)
}

func (s *OrderUnitUsecase) GetOrderUnitById(id uint) dto.Response {

	uri := getOrderUnitUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return s.GetRequest.Main(uri)
}

package usecase

import (
	"core/dto"
	request "core/request"
	"os"
	"strconv"
)

type OrderArticleUsecaseContract interface {
	GetOrderArticles() dto.Response
	GetOrderArticleByID(uint) dto.Response
}

type OrderArticleUsecase struct {
	GetRequest request.GetRequest
}

func ProviderOrderArticleUsecase(g request.GetRequest) OrderArticleUsecase {
	return OrderArticleUsecase{
		GetRequest: g,
	}
}

func getOrderArticleUri() string {

	uri := os.Getenv("ORDER_URI")
	uri += "/order/article"

	return uri
}

// Implementation

func (oa *OrderArticleUsecase) GetOrderArticles() dto.Response {

	uri := getOrderArticleUri()

	return oa.GetRequest.Main(uri)
}

func (oa *OrderArticleUsecase) GetOrderArticleByID(id uint) dto.Response {

	uri := getOrderArticleUri()
	uri += "/" + strconv.FormatUint(uint64(id), 10)

	return oa.GetRequest.Main(uri)
}

package controller

import (
	"fmt"
	"net/http"
	"service/config"
	"service/dto"
	"service/usecase"

	"github.com/labstack/echo"
)

type ServiceController struct {
	ServiceUsecase usecase.ServiceUsecase
}

func ProviderServiceController(s usecase.ServiceUsecase) ServiceController {
	return ServiceController{
		ServiceUsecase: s,
	}
}

func (s *ServiceController) Create(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		fmt.Println(e.Bind(&service))
		fmt.Println(service.Name)
		fmt.Println(service.Description)
		fmt.Println(service.Price)
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := s.ServiceUsecase.Create(service)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, config.CreateServiceSuccess)
}

package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type ServiceController struct {
	ServiceUsecase usecase.ServiceUsecase
}

func ProviderServiceController(s usecase.ServiceUsecase) ServiceController {
	return ServiceController{
		ServiceUsecase: s,
	}
}

func (s *ServiceController) GetServices(e echo.Context) error {

	res := s.ServiceUsecase.GetServices()

	return CheckResponse(e, res)
}

func (s *ServiceController) GetServiceById(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := s.ServiceUsecase.GetServiceById(uint(id))

	return CheckResponse(e, res)
}

func (s *ServiceController) Create(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res := s.ServiceUsecase.CreateService(request)

	return CheckResponse(e, res)
}

func (s *ServiceController) Update(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res := s.ServiceUsecase.UpdateService(request)

	return CheckResponse(e, res)

}

func (s *ServiceController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := s.ServiceUsecase.DeleteService(uint(id))

	return CheckResponse(e, res)
}

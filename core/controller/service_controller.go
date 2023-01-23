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

	filter := e.QueryParam("filter")
	status := e.QueryParam("status")

	res := s.ServiceUsecase.GetServices(filter, status)

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

func (s *ServiceController) CreateService(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res := s.ServiceUsecase.CreateService(request)

	return CheckResponse(e, res)
}

func (s *ServiceController) UpdateService(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res := s.ServiceUsecase.UpdateService(request)

	return CheckResponse(e, res)

}

func (s *ServiceController) DeleteService(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := s.ServiceUsecase.DeleteService(uint(id))

	return CheckResponse(e, res)
}

func (s *ServiceController) ActivateService(e echo.Context) error {

	status := e.QueryParam("Status")
	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	res := s.ServiceUsecase.ActivateService(uint(id), status)

	return CheckResponse(e, res)
}

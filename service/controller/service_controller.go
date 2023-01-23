package controller

import (
	"net/http"
	"service/config"
	"service/dto"
	"service/usecase"
	"strconv"

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

func (s *ServiceController) GetAll(e echo.Context) error {

	filter := e.QueryParam("filter")
	status := e.QueryParam("status")

	services := s.ServiceUsecase.GetAll(filter, status)

	if len(services) == 0 {
		return config.SuccessResponse(e, nil, config.ServiceNotFound)
	}

	return config.SuccessResponse(e, services, config.GetServiceSuccess)
}

func (s *ServiceController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	service := s.ServiceUsecase.GetByID(uint(id))

	if service.ID == 0 {
		return config.SuccessResponse(e, nil, config.ServiceNotFound)
	}

	return config.SuccessResponse(e, service, config.GetServiceSuccess)
}

func (s *ServiceController) Create(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := s.ServiceUsecase.Create(service)

	return CheckResponse(e, err, config.CreateServiceSuccess)
}

func (s *ServiceController) Update(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := s.ServiceUsecase.Update(service)

	return CheckResponse(e, err, config.UpdateServiceSuccess)
}

func (s *ServiceController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = s.ServiceUsecase.Delete(uint(id))

	return CheckResponse(e, err, config.DeleteServiceSuccess)
}

func (s *ServiceController) Activate(e echo.Context) error {

	status := e.QueryParam("Status")
	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = s.ServiceUsecase.Activate(uint(id), status)

	return CheckResponse(e, err, config.ActivateServiceSuccess)
}

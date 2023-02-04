package controller

import (
	"net/http"
	"service/config"
	"service/dto"
	"service/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
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

	var query_param dto.QueryParam

	if e.Bind(&query_param) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	services := s.ServiceUsecase.GetAll(query_param)

	if len(services) == 0 {
		return SuccessResponse(e, nil, config.ServiceNotFound)
	}

	return SuccessResponse(e, services, config.GetServiceSuccess)
}

func (s *ServiceController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	service := s.ServiceUsecase.GetByID(uint(id))

	if service.ID == 0 {
		return SuccessResponse(e, nil, config.ServiceNotFound)
	}

	return SuccessResponse(e, service, config.GetServiceSuccess)
}

func (s *ServiceController) Create(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	service_dto, err, err_code := s.ServiceUsecase.Create(service)

	return CheckCreateResponse(e, service_dto, err, err_code, config.CreateServiceSuccess)
}

func (s *ServiceController) Update(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err, err_code := s.ServiceUsecase.Update(service)

	return CheckResponse(e, err, err_code, config.UpdateServiceSuccess)
}

func (s *ServiceController) Delete(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := s.ServiceUsecase.Delete(service)

	return CheckResponse(e, err, err_code, config.DeleteServiceSuccess)
}

func (s *ServiceController) Activate(e echo.Context) error {

	var service dto.Service

	if e.Bind(&service) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := s.ServiceUsecase.Activate(service)

	if !service.IsActived {
		return CheckResponse(e, err, err_code, config.DeactivateServiceSuccess)
	}

	return CheckResponse(e, err, err_code, config.ActivateServiceSuccess)
}

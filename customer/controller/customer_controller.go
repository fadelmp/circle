package controller

import (
	"customer/config"
	"customer/dto"
	"customer/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

type CustomerController struct {
	CustomerUsecase usecase.CustomerUsecase
}

func ProviderCustomerController(c usecase.CustomerUsecase) CustomerController {
	return CustomerController{CustomerUsecase: c}
}

func (c *CustomerController) GetAll(e echo.Context) error {

	var query_param dto.QueryParam

	if e.Bind(&query_param) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	customers := c.CustomerUsecase.GetAll(query_param)

	if len(customers) == 0 {
		return SuccessResponse(e, nil, config.CustomerNotFound)
	}

	return SuccessResponse(e, customers, config.GetCustomerSuccess)
}

func (c *CustomerController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	customer := c.CustomerUsecase.GetByID(uint(id))

	if customer.ID == 0 {
		return SuccessResponse(e, nil, config.CustomerNotFound)
	}

	return SuccessResponse(e, customer, config.GetCustomerSuccess)
}

func (c *CustomerController) Create(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err, err_code := c.CustomerUsecase.Create(customer)

	return CheckResponse(e, err, err_code, config.CreateCustomerSuccess)
}

func (c *CustomerController) Update(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	err, err_code := c.CustomerUsecase.Update(customer)

	return CheckResponse(e, err, err_code, config.UpdateCustomerSuccess)
}

func (c *CustomerController) Delete(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := c.CustomerUsecase.Delete(customer)

	return CheckResponse(e, err, err_code, config.DeleteCustomerSuccess)
}

func (c *CustomerController) Activate(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	err, err_code := c.CustomerUsecase.Activate(customer)

	if !customer.IsActived {
		return CheckResponse(e, err, err_code, config.DeactivateCustomerSuccess)
	}

	return CheckResponse(e, err, err_code, config.ActivateCustomerSuccess)
}

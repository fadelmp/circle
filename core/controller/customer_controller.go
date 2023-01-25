package controller

import (
	"core/config"
	"core/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type CustomerController struct {
	CustomerUsecase usecase.CustomerUsecase
}

func ProviderCustomerController(c usecase.CustomerUsecase) CustomerController {
	return CustomerController{CustomerUsecase: c}
}

func (c *CustomerController) GetCustomers(e echo.Context) error {

	filter := e.QueryParam("filter")
	status := e.QueryParam("status")

	res := c.CustomerUsecase.GetCustomers(filter, status)

	return CheckResponse(e, res)
}

func (c *CustomerController) GetCustomerById(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := c.CustomerUsecase.GetCustomerById(uint(id))

	return CheckResponse(e, res)
}

func (c *CustomerController) CreateCustomer(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := c.CustomerUsecase.CreateCustomer(request)

	return CheckResponse(e, res)
}

func (c *CustomerController) UpdateCustomer(e echo.Context) error {

	var request interface{}

	if e.Bind(&request) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, 3, config.BadRequest)
	}

	res := c.CustomerUsecase.UpdateCustomer(request)

	return CheckResponse(e, res)

}

func (c *CustomerController) DeleteCustomer(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := c.CustomerUsecase.DeleteCustomer(uint(id))

	return CheckResponse(e, res)
}

func (c *CustomerController) ActivateCustomer(e echo.Context) error {

	status := e.Param("status")
	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, 3, config.BadRequest)
	}

	res := c.CustomerUsecase.ActivateCustomer(uint(id), status)

	return CheckResponse(e, res)
}

package controller

import (
	"customer/config"
	"customer/dto"
	"customer/service"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func ProviderCustomerController(c service.CustomerService) CustomerController {
	return CustomerController{CustomerService: c}
}

func (c *CustomerController) Create(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := c.CustomerService.Create(customer)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.CreateCustomerSuccess)
}

package controller

import (
	"customer/config"
	"customer/dto"
	"customer/usecase"
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

func (c *CustomerController) GetAll(e echo.Context) error {

	customers := c.CustomerUsecase.GetAll()

	if len(customers) == 0 {
		return config.SuccessResponse(e, nil, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, customers, config.GetCustomerSuccess)
}

func (c *CustomerController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	customer := c.CustomerUsecase.GetByID(uint(id))

	if customer.ID == 0 {
		return config.SuccessResponse(e, nil, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, customer, config.GetCustomerSuccess)
}

func (c *CustomerController) Create(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := c.CustomerUsecase.Create(customer)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.CreateCustomerSuccess)
}

func (c *CustomerController) Update(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err := c.CustomerUsecase.Update(customer)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.UpdateCustomerSuccess)
}

func (c *CustomerController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = c.CustomerUsecase.Delete(uint(id))

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.DeleteCustomerSuccess)
}

func (c *CustomerController) Activate(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = c.CustomerUsecase.ActiveStatus(uint(id), true)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.ActivateCustomerSuccess)
}

func (c *CustomerController) Deactivate(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("ID"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusBadRequest, config.BadRequest)
	}

	err = c.CustomerUsecase.ActiveStatus(uint(id), false)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, nil, config.DeactivateCustomerSuccess)
}

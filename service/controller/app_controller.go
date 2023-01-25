package controller

import (
	"net/http"
	"service/config"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

func CheckResponse(e echo.Context, err error, err_code int, message string) error {

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err_code, err.Error())
	}

	return config.SuccessResponse(e, nil, message)
}

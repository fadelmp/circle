package controller

import (
	"core/config"
	"core/dto"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

func CheckResponse(e echo.Context, response dto.Response) error {

	if response.ErrorMessage != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, response.Result.ErrorCode, response.ErrorMessage)
	}

	if response.ResponseCode != 200 {
		return config.ErrorResponse(e, response.ResponseCode, response.Result.ErrorCode, response.Result.Messages)
	}

	return config.SuccessResponse(e, response.Result)
}

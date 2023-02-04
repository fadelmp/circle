package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

type SuccessResp struct {
	Messages interface{}
	Data     interface{}
}

type ErrorResp struct {
	Messages  interface{}
	ErrorCode int
}

func SuccessResponse(c echo.Context, data interface{}, message interface{}) error {

	success_response := 200

	resp := &SuccessResp{
		Messages: message,
		Data:     data,
	}

	c.Response().WriteHeader(success_response)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSONPretty(success_response, resp, "  ")
}

func ErrorResponse(c echo.Context, httpErrorCode int, error_code int, messages interface{}) error {

	resp := &ErrorResp{
		Messages:  messages,
		ErrorCode: error_code,
	}

	c.Response().WriteHeader(httpErrorCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSONPretty(httpErrorCode, resp, "  ")
}

func CheckResponse(
	e echo.Context,
	err error,
	err_code int,
	message string,
) error {

	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err_code, err.Error())
	}

	return SuccessResponse(e, nil, message)
}

func CheckCreateResponse(
	e echo.Context,
	data interface{},
	err error,
	err_code int,
	message string,
) error {

	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err_code, err.Error())
	}

	return SuccessResponse(e, data, message)
}

package config

import (
	"core/dto"

	"github.com/labstack/echo"
)

type SuccessResp struct {
	Messages interface{}
	Data     interface{}
}

type ErrorResp struct {
	Messages interface{}
}

func SuccessResponse(c echo.Context, result dto.Result) error {
	resp := &SuccessResp{
		Messages: result.Messages,
		Data:     result.Data,
	}
	c.Response().WriteHeader(200)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(200, resp, "  ")
}

func ErrorResponse(c echo.Context, errorCode int, Messages interface{}) error {
	resp := &ErrorResp{
		Messages: Messages,
	}
	c.Response().WriteHeader(errorCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(errorCode, resp, "  ")
}

package config

import "github.com/labstack/echo"

type SuccessResp struct {
	Messages interface{}
	Data     interface{}
}

type ErrorResp struct {
	Messages interface{}
}

func SuccessResponse(c echo.Context, data interface{}, message interface{}) error {
	resp := &SuccessResp{
		Messages: message,
		Data:     data,
	}
	c.Response().WriteHeader(200)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(200, resp, "  ")
}

func ErrorResponse(c echo.Context, errorCode int, messages interface{}) error {
	resp := &ErrorResp{
		Messages: messages,
	}
	c.Response().WriteHeader(errorCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(errorCode, resp, "  ")
}

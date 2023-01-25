package config

import "github.com/labstack/echo"

type SuccessResp struct {
	Messages interface{}
	Data     interface{}
}

type ErrorResp struct {
	Messages  interface{}
	ErrorCode int
}

func SuccessResponse(c echo.Context, data interface{}, message interface{}) error {
	resp := &SuccessResp{
		Messages: message,
		Data:     data,
	}
	c.Response().WriteHeader(200)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSONPretty(200, resp, "  ")
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

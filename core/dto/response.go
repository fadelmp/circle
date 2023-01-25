package dto

type Response struct {
	ResponseCode int
	ErrorMessage error
	Result       Result
}

type Result struct {
	ErrorCode interface{}
	Messages  interface{}
	Data      interface{}
}

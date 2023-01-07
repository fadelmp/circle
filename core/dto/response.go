package dto

type Response struct {
	ResponseCode int
	ErrorMessage error
	Result       Result
}

type Result struct {
	Messages interface{}
	Data     interface{}
}

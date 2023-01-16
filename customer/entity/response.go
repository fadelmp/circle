package entity

type Response struct {
	ResponseCode int
	ErrorMessage error
	Result       Result
}

type Result struct {
	Messages interface{}
	Data     Data
}

type Data struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
